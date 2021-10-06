// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a ReplicatedBucket component resource.
type ReplicatedBucketArgs struct {
	// Destination region for the replicated bucket.
	DestinationRegion pulumi.StringInput `pulumi:"destinationRegion"`
}

// The ReplicatedBucket component resource.
type ReplicatedBucket struct {
	pulumi.ResourceState

	SourceBucket      *s3.Bucket `pulumi:"sourceBucket"`
	DestinationBucket *s3.Bucket `pulumi:"destinationBucket"`
}

// NewReplicatedBucket creates a new ReplicatedBucket component resource.
func NewReplicatedBucket(ctx *pulumi.Context,
	name string, args *ReplicatedBucketArgs, opts ...pulumi.ResourceOption) (*ReplicatedBucket, error) {
	if args == nil {
		args = &ReplicatedBucketArgs{}
	}

	component := &ReplicatedBucket{}
	err := ctx.RegisterComponentResource("aws-s3-replicated-bucket:index:ReplicatedBucket", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// Create a provider for the destination region.
	provider, err := aws.NewProvider(ctx, fmt.Sprintf("%sProvider", name), &aws.ProviderArgs{
		Region: args.DestinationRegion,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Create the destination bucket.
	dst, err := s3.NewBucket(ctx, fmt.Sprintf("%sDestination", name), &s3.BucketArgs{
		Versioning: &s3.BucketVersioningArgs{
			Enabled: pulumi.BoolPtr(true),
		},
	}, pulumi.Parent(component),
		pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	replicationAssumeRolePolicy, _ := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Action": "sts:AssumeRole",
				"Effect": "Allow",
				"Sid":    "",
				"Principal": map[string]interface{}{
					"Service": []string{
						"s3.amazonaws.com",
					},
				},
			},
		},
	})
	// Create a role for replication.
	replicationRole, err := iam.NewRole(ctx, fmt.Sprintf("%sReplicationRole", name), &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(string(replicationAssumeRolePolicy)),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Create the source bucket with replication configuration.
	src, err := s3.NewBucket(ctx, fmt.Sprintf("%sSource", name), &s3.BucketArgs{
		Versioning: &s3.BucketVersioningArgs{
			Enabled: pulumi.BoolPtr(true),
		},
		ReplicationConfiguration: &s3.BucketReplicationConfigurationArgs{
			Role: replicationRole.Arn,
			Rules: &s3.BucketReplicationConfigurationRuleArray{
				&s3.BucketReplicationConfigurationRuleArgs{
					Destination: &s3.BucketReplicationConfigurationRuleDestinationArgs{
						Bucket: dst.Arn,
					},
					Status:   pulumi.String("Enabled"),
					Priority: pulumi.Int(1),
				},
			},
		},
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Create the replication policy.
	replicationPolicyString := pulumi.All(src.Arn, dst.Arn).ApplyT(
		func(args []interface{}) string {
			sourceBucketArn := args[0].(string)
			destinationBucketArn := args[1].(string)
			b, _ := json.Marshal(map[string]interface{}{
				"Version": "2012-10-17",
				"Statement": []map[string]interface{}{
					{
						"Effect": "Allow",
						"Action": []string{
							"s3:GetObjectVersionForReplication",
							"s3:GetObjectVersionAcl",
						},
						"Resource": []string{
							sourceBucketArn + "/*",
						},
					},
					{
						"Effect": "Allow",
						"Action": []string{
							"s3:ListBucket",
							"s3:GetReplicationConfiguration",
						},
						"Resource": sourceBucketArn,
					},
					{
						"Effect": "Allow",
						"Action": []string{
							"s3:ReplicateObject",
							"s3:ReplicateDelete",
							"s3:ReplicateTags",
							"s3:GetObjectVersionTagging",
						},
						"Resource": destinationBucketArn + "/*",
					},
				},
			})

			return string(b)
		},
	)
	replicationPolicy, err := iam.NewPolicy(ctx, fmt.Sprintf("%sReplicationPolicy", name), &iam.PolicyArgs{
		Policy: replicationPolicyString,
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// Attach the policy to our replication role.
	_, err = iam.NewPolicyAttachment(ctx, fmt.Sprintf("%sReplicationPolicyAttachment", name), &iam.PolicyAttachmentArgs{
		Roles:     pulumi.ToArrayOutput([]pulumi.Output{replicationRole.Name}),
		PolicyArn: replicationPolicy.Arn,
	}, pulumi.Parent(component))

	component.DestinationBucket = dst
	component.SourceBucket = src

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
		"destinationBucket": dst,
		"sourceBucket":      src,
	}); err != nil {
		return nil, err
	}

	return component, nil
}
