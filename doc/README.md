# Table of contents

1. [Description](#Description)
2. [Why Mordo Image Service](#Why-Mordo-Image-Service)
3. [Supported Opearations](#Supported-Opearations)
4. [Platforms](#Platforms)
5. [Usage](#Usage)
    * [Routes](#Routes)
    * [Query Strings](#Query-Strings)
6. [Service Configuration](#Service-Configuration)
7. [Service Setup](#Service-Setup)
    * [AWS](setup_AWS.md)

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
4. Smart Crop
5. Watermark
6. Sharpen

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

For testing purposes `/process` AWS Api Gateway endpoint is available. This endpoint won't cache resulting image at theedge locations.

### Query Strings

Image operations are defined by query string passed to endpoint.

1. `width`/`heigth` **Resize** - use one to define new size of the image. For example:

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?width=1000`
    * `https://demo.mordo.io/origin_2k.jpg?heigth=1000`

2. `dpr` - set Dual Pixel Ratio, number. It is kind of resize, but act like scale factor. Vaules: `0` - `infinity`.

    Usage examples:
      * `https://demo.mordo.io/origin_2k.jpg?dpr=1.5` - DPR (scale) factor 1.5, which means image will be 150% size of original

    _Note: Please avoid using with **Resize** operation, cause under the hood **DPR** and **Resize** is almost same operation and modifies image in same way_

    _Note: If `dpr` > 1 image will be enlarged. Enlargin is heavy opertion_

3. `c_top`, `c_left`, `c_x`, `c_y` **Smart Crop** - used to define `Top`, `Left` position and `X`, `Y` axes of **Smart Crop** operation

    `c_top`, `c_left` - top and left postition to crop.

    `c_x`, `c_y` - size of picture after crop.

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?c_top=500&c_left=700&c_x=400&c_y=700` - crop 400x700 image
    * `https://demo.mordo.io/origin_2k.jpg?width=1000&c_top=500&c_left=700&c_x=400&c_y=700` - resize image to 1000 px by **width** and crop to 400x700

4. `sc_x`, `sc_y` **Smart Crop** - used to define `X`, `Y` axes of **Smart Crop** operation

    `sc_x`, `sc_y` - size of picture after crop.

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?sc_x=400&sc_y=700` - crop 400x700 image
    * `https://demo.mordo.io/origin_2k.jpg?width=1000&sc_x=400&sc_y=700` - resize image to 1000 px by **width** and crop to 400x700

5. `w_x`, `w_y`, `w_s` **Watermark** - use to define watermark position and watermark scale.

    `w_x` - position by X axis, string, required parameter. Available values:
    * `right`
    * `left`
    * `center`

    `w_y` - position by Y axis, string, required parameter. Available values:
    * `top`
    * `bottom`
    * `center`

      `w_s` - watermarkscale factor, number, optional paramter. Vaules: `0` - `infinity`.

      _Note: setting `w_s` value more than `1` will have severe impact on performance._

    Usage examples:
    * `https://demo.mordo.io/origin_2k.jpg?w_x=left&w_y=bottom` - watermark to left bottom corner
    * `https://demo.mordo.io/origin_2k.jpg?w_x=right&w_y=center&w_s=0.5` - watermark to right by X, center by Y
    * `https://demo.mordo.io/origin_2k.jpg?width=1000&w_x=right&w_y=center` - resize image to 1000 px by **width** watermark to right by **X**, center by **Y**, scale watermark to **50%** of original size.

6. `sharpen` - **Sharpen operation** - use to apply sharpening to image. Vaules: `t` - apply sharpening.

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

* [AWS](setup_AWS.md)
