import * as s3 from "@pulumi/aws-s3-replicated-bucket";

const bucket = new s3.Bucket("bucket", {
    destinationRegion: "us-east-1",
});

export const src = bucket.sourceBucket;
export const dst = bucket.destinationBucket;
