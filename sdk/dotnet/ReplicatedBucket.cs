// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsS3ReplicatedBucket
{
    [AwsS3ReplicatedBucketResourceType("aws-s3-replicated-bucket:index:ReplicatedBucket")]
    public partial class ReplicatedBucket : Pulumi.ComponentResource
    {
        /// <summary>
        /// Bucket to which data should be replicated.
        /// </summary>
        [Output("destinationBucket")]
        public Output<Pulumi.Aws.S3.Bucket> DestinationBucket { get; private set; } = null!;

        /// <summary>
        /// test stuff
        /// </summary>
        [Output("locationPolicy")]
        public Output<Pulumi.Kubernetes.Gcp/gke.Outputs.NodePoolAutoscaling?> LocationPolicy { get; private set; } = null!;

        /// <summary>
        /// Bucket to which objects are written.
        /// </summary>
        [Output("sourceBucket")]
        public Output<Pulumi.Aws.S3.Bucket> SourceBucket { get; private set; } = null!;


        /// <summary>
        /// Create a ReplicatedBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ReplicatedBucket(string name, ReplicatedBucketArgs args, ComponentResourceOptions? options = null)
            : base("aws-s3-replicated-bucket:index:ReplicatedBucket", name, args ?? new ReplicatedBucketArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ReplicatedBucketArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Region to which data should be replicated.
        /// </summary>
        [Input("destinationRegion", required: true)]
        public Input<string> DestinationRegion { get; set; } = null!;

        public ReplicatedBucketArgs()
        {
        }
    }
}
