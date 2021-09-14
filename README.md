# Replicated Bucket Pulumi Component Provider 

Create a [s3](https://www.pulumi.com/docs/reference/pkg/aws/s3/bucket/) that replicates its contents to another [region](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Concepts.RegionsAndAvailabilityZones.html). 

--- 

## Installing
This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/replicatedbucket

or `yarn`:

    $ yarn add @pulumi/replicatedbucket

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_replicatedbucket

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-replicatedbucket/sdk

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.ReplicatedBucket

## Concept

---

The Pulumi replicated Bucket provides a simple and correct implementation of a
s3 bucket replicated to another region. 

``` typescript
const bucket = new replicatedBucket.Bucket("bucket", {
    destinationRegion: "us-east-1",
});
```

``` c#
var bucket = new ReplicatedBucket.Bucket(ctx, "bucket", new ReplicatedBucket.BucketArgs{
    DestinationRegion: pulumi.String("us-east-1")
})
```

``` go
bucket, err = replicatedbucket.NewBucket(ctx, "bucket", &replicatedbucket.BucketArgs{
	DestinationRegion: pulumi.String("us-east-1"),
})
```

``` python
bucket = pulumi_replicatedbucket.Bucket("bucket", destination_region="us-east-1")
```

