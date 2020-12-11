# Table of contents

- [Table of contents](#table-of-contents)
  - [Description](#description)
  - [Why Mordo Image Service](#why-mordo-image-service)
  - [Supported Opearations](#supported-opearations)
  - [Platforms](#platforms)
  - [Usage](#usage)
    - [Routes](#routes)
    - [Query Strings](#query-strings)
  - [Service Configuration](#service-configuration)
  - [Service Setup](#service-setup)
  - [Local Testing](#local-testing)
    - [Config](#config)
    - [Run](#run)

## Description

Mordo image service is designed with serverless im mind.
This service is exposed as HTTP endpoint and used TBD

## Why Mordo Image Service

Mordo Image Service is designed with serverless in the soul. Main advantage of serverless applications is mintenance costs. Serverless services require zero infrastructure maintenance and configuration. Second big advantage is running costs. With serverless "you pay as you go", you pay only for time that service works (invocation time) which generally much cheaper than traditional "VM running service".  

## Supported Opearations

Next image operations are supported:

1. Resize (with preserving aspect ratio)
2. DPR (Dual Pixel Ratio)
3. Crop
4. Watermark
5. Sharpen

## Platforms

Supported:

1. _AWS_ (AWS Lambda)

Future:

1. _GCP_
2. _Standalone server_

## Usage

### Routes

Mordo API has only one endpoint `/{image path}` exposed via:

* _CloudFront_ for _AWS_

For testing purposes `/process` AWS Api Gateway endpoint is available. This endpoint won't cache resulting image at the edge locations.

### Query Strings

Image operations are defined by query string passed to endpoint.

1. `width`/`height` **Resize** - use one to define new size of the image. For example:

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?width=1000`
    * `https://demo.mordo.io/origin_2k.jpg?height=1000`

2. `dpr` - set Dual Pixel Ratio, number. It is kind of resize, but act like scale factor. Vaules: `0` - `infinity`.

    Usage examples:
      * `https://demo.mordo.io/origin_2k.jpg?dpr=1.5` - DPR (scale) factor 1.5, which means image will be 150% size of original

    _Note: Please avoid using with **Resize** operation, cause under the hood **DPR** and **Resize** is almost same operation and modifies image in same way_

    _Note: If `dpr` > 1 image will be enlarged. Enlarging is heavy operation_

3. `w_x`, `w_y`, `w_s` **Watermark** - use to define watermark position and watermark scale.

    `w_x` - position by X axis, string, required parameter. Available values:
    * `right`
    * `left`
    * `center`

    `w_y` - position by Y axis, string, required parameter. Available values:
    * `top`
    * `bottom`
    * `center`

      `w_s` - watermark scale factor, number, optional parameter. Values: `0` - `infinity`.

      _Note: setting `w_s` value more than `1` will have severe impact on performance._

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?w_x=left&w_y=bottom` - watermark to left bottom corner
    * `https://demo.mordo.io/origin_2k.jpg?w_x=right&w_y=center&w_s=0.5` - watermark to right by X, center by Y
    * `https://demo.mordo.io/origin_2k.jpg?width=1000&w_x=right&w_y=center` - resize image to 1000 px by **width** watermark to right by **X**, center by **Y**, scale watermark to **50%** of original size.

4. `sharpen` - **Sharpen operation** - use to apply sharpening to image. Vaules: `t` - apply sharpening.

    Predefined sharpening params:
    * `radius`: 1 - Gaussian Blur radius for high-frequency signal
    * `X1`: 2 - flat/jaggy threshold
    * `Y2`: 5 - maximum amount of brightening
    * `Y3`: 10 - maximum amount of darkening
    * `M1`: 0 - slope for flat areas
    * `M2`: 3 - slope for jaggy areas

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?sharpen=t` - apply sharpening.

## Service Configuration

Mordo Image service is configured via environment variables.

* `"DEBUG":                   "true"` - Mordo Debug mode. More logging, no S3 cache, CloudFront caching still applies.
* `"PLATFORM":                "AWS"`
* `"AWS_S3_BUCKET_PROCESSED": "foo-bar"`
* `"WATERMARK":               "true"`
* `"WATERMARK_PATH":          "foo/bar/baz/watermark"`
* `"WATERMARK_SCALE":         "0.15"`
* `"IMAGE_QUALITY":           "90"`
* `"IMAGE_LOSSLESS":          "true"`
* `"IMAGE_COMPRESSION":       "0"`
* `"IMAGE_ENLARGE":           "TRUE"`

## Service Setup

The best way to setup Mordo is [Terraform](https://www.terraform.io/downloads.html)

Terraform install resources is located at `./terraform`.

* [AWS](doc/setup_AWS.md)

## Local Testing

### Config

In order to run locally you need to create S3 bucket for resulting images and set `AWS_S3_BUCKET_PROCESSED` in `.env`.

Also You need to fill `.envlocal.sample` with `AWS_ACCESS_KEY` and `AWS_SECRET_KEY`. This user must have next permissions:

```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "1",
            "Effect": "Allow",
            "Action": [
                "s3:PutObjectTagging",
                "s3:PutObject",
                "s3:PutObjectAcl",
                "s3:GetObject",
                "s3:ListBucket",
                "s3:ListAllMyBuckets",
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::AWS_S3_BUCKET_PROCESSED/*",
                "arn:aws:s3:::AWS_S3_BUCKET_PROCESSED"
            ]
        },
        {
            "Sid": "2",
            "Effect": "Allow",
            "Action": [
                "s3:GetObject",
                "s3:ListBucket",
                "s3:ListAllMyBuckets",
                "s3:GetBucketLocation"
            ],
            "Resource": [
                "arn:aws:s3:::AWS_S3_BUCKET_WITH_IMAGES/*",
                "arn:aws:s3:::AWS_S3_BUCKET_WITH_IMAGES"
            ]
        },
        {
            "Sid": "3",
            "Effect": "Allow",
            "Action": [
                "s3:ListAllMyBuckets",
                "s3:GetBucketLocation"
            ],
            "Resource": "*"
        }
    ]
}
```

Modify next blocks of `doc/apigw_event/blank.json`

_Note: You can `doc/apigw_event/*` are events that **AWS API Gateway** passes to **AWS Lambda**. Thay are passed to local container which emulates **AWS Lambda** behaviour._

```
  "queryStringParameters": {
    "width": "1000"
  },
```

and

```
  "stageVariables": {
    "baz": "qux",
    "s3_bucket": "AWS_S3_BUCKET_WITH_IMAGES",
    "s3_bucket_region": "us-west-1"
  },
```

### Run

```
# sh scripts/build_linux.sh aws 1>/dev/null

# cat doc/apigw_event/blank.json |  docker run --rm -i -v "$PWD"/pkg:/var/task --env-file=.envlocal --env-file=.env lambci/lambda:go1.x mordo 2>&1
```

or

```
sh scripts/run_aws_local.sh
```
