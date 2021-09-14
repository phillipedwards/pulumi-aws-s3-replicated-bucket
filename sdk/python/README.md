# Replicated Bucket Pulumi Component Provider 

Create a [s3](https://www.pulumi.com/docs/reference/pkg/aws/s3/bucket/) that replicates its contents to another [region](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html). 

--- 

This program is written in Go, but is consumable in any language Pulumi supports. 

## Example
``` typescript
const bucket = new replicatedBucket.Bucket("bucket", {
    destinationRegion: "us-east-1",
});

export const src = bucket.sourceBucket;
export const dst = bucket.destinationBucket;
```
