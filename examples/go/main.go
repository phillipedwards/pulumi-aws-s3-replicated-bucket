package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
	"github.com/pulumi/pulumi-aws-s3-replicated-bucket/sdk/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		bucket, err := aws.NewReplicatedBucket(ctx, "my-bucket", &aws.ReplicatedBucketArgs{
			DestinationRegion: pulumi.String("us-east-1"),
		})
		if err != nil {
			return err
		}

		id := func(b *s3.Bucket) pulumi.IDOutput { return b.ID() }

		ctx.Export("src", bucket.SourceBucket.ApplyT(id))
		ctx.Export("dst", bucket.DestinationBucket.ApplyT(id))
		return nil
	})
}
