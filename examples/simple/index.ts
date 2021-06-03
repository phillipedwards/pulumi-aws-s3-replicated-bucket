import * as replicatedBucket from "@pulumi/replicatedbucket";

const bucket = new replicatedBucket.Bucket("bucket", {
    destinationRegion: "us-east-1",
});

export const src = bucket.sourceBucket;
export const dst = bucket.destinationBucket;
