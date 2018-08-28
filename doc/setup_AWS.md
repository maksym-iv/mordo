# Description

The best way to setup this service in **AWS** is use [Terraform](https://www.terraform.io/downloads.html). Provisioning resources located at `./terraform/aws`.
Terraform won't delete any reources, all IAM roles created wo't have write access to any existing resources except CloudWatch, which is required forlogging.

_Note: Pay attention to keep/backup `terraform.tfstate` file that is created during provisioning. This file saves state and required for resource modifications_

## Setup

### Configuration

For propper setup next variables configured. File with Terraform sample variables could be found here `terraform/aws/_vars.auto.tfvars.sample`.

Variables description with recommended values:

* `aws_access_key = ""` - AWS Access Key for service set up.
* `aws_secret_key = ""` - AWS Secret Key for service set up.
* `aws_region = ""` - AWS Region for service set up.

* `prefix = "dev"` - Prefix to add for created resources

* `s3_buckets = ["bucket1", "bucket2"]` - Buckets with source images.

  _Note: Terraform will create CloudFront Distribution and API Gatevays per each bucket_

* `s3_bucket_deploy = "mordo-deploy-bucket"` - Bucket to store AWS Lambda code. Used for Lambda deployment
* `s3_bucket_prc = "bucket-images-processed"` - Bucket to store processed images.
* `s3_bucket_prc_lifecycle = "true"` - Enable or disable lifecycle rules on bucket with processed images. Values: true/false
* `s3_bucket_prc_lifecycle_expire = "3"` - If lifecycle is enabled on on bucket with processed images, sets expiration time in days. This will also set CloudFron MaxTTL for cache.

* `lambda_concurrency = 5` - concurent executions of the service
* `lambda_timeout = 10` - AWS Lambda execution timeout. Set this according to max size of your images.
* `lambda_memory = 512` - AWS Lambda memory limit. Set this according to max size of your images.
  _Note: More mem = faster CPU. Faster CPU = less time for image processing. Noticable performance increase will stop at ~1500 mb_
* `lambda_DEBUG = "true"` - Mordo Debug mode. More logging, no S3 cache, CloudFront caching still applies.
* `lambda_WATERMARK = "true"` - Enable watermarking. true/false
* `lambda_WATERMARK_PATH = "watermark.png"` - Watermark path if lambda_WATERMARK enabled.
* `lambda_WATERMARK_SCALE = "0.3"` - Watermark scale according to image size, if lambda_WATERMARK enabled.
* `lambda_IMAGE_QUALITY = "90"` - JPEG Image quality. 0-100
* `lambda_IMAGE_LOSSLESS = "true"` - WEBP Image quality. true/false
* `lambda_IMAGE_COMPRESSION = "0"` - PNG Image compression. 0-10
* `lambda_IMAGE_ENLARGE = "false"` - Allow image enlarge during resize. true/false

* `apigw_throttling_rate_limit = 10` - AWS API Gateway throttling rate limit
* `apigw_throttling_burst_limit = 5` - AWS API Gateway throttling burst limit
  _Deatail on AWS API Gateway throttling can be found [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-request-throttling.html)_

* `cf_price_class = "PriceClass_200"` - CloudFront price class. Details can be found [here](https://aws.amazon.com/cloudfront/pricing/)
* `cf_min_ttl = 0` - The minimum amount of time (in seconds) that you want objects to stay in CloudFront caches before CloudFront queries your origin to see whether the object has been updated
* `cf_default_ttl = 87000` - The default amount of time (in seconds) that an object is in a CloudFront cache before CloudFront forwards another request in the absence of an Cache-Control max-age or Expires header

### Setup IAM

For setup to work itis required to create IAM user with next permissions

_Note: S3 permissions (**Sid 4**) should be set according to your `prefix`, `s3_bucket_prc`, `s3_bucket_deploy`_

Permissions [IAM Doc](setup_IAM.json)

* `${prefix}` - prefix for the created S3 buckets (must be same with `pefix` in Terraform variable file).
* `${s3_bucket_prc}` - bucket to store processed images (must be same with `s3_bucket_prc` in Terraform variable file).
* `${s3_bucket_deploy}` - bucket to store AWS Lambda code (must be same with `s3_bucket_deploy` in Terraform variable file).

### Run

After filling Terraform sample variables copy it to `terraform/aws/_vars.auto.tfvars` and run

```shell
terraform init
terraform apply
```

Terraform will create next resources

* S3 Bucket for AWS La,bda deployment. Will be cleaned every day by lifecycle policy.
* S3 Bucket for processed images
* AWS API Gateway (1 Gateway per each bucket from `s3_buckets`)
* AWS CloudFront (1 Gateway per each bucket from `s3_buckets`)
* AWS IAM Role for Lambda function with read permissions to each bucket from `s3_buckets`
* AWS Lambda Fuction
* AWS Lambda permission for AWS API Gateway