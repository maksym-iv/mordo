# AWS credentials
variable "aws_access_key" {
  type        = "string"
  description = ""
}

variable "aws_secret_key" {
  type        = "string"
  description = ""
}

variable "aws_region" {
  type        = "string"
  description = ""
}

# END AWS credentials

# Misc
variable "prefix" {
  type        = "string"
  description = "Resource name prefix"
}

# S3 config
variable "s3_buckets" {
  type        = "list"
  description = "Source S3 buckets"
}

variable "s3_bucket_deploy" {
  type        = "string"
  description = "Bucket to store AWS Lambda code. Used for Lambda deployment"
}

variable "s3_bucket_prc" {
  type        = "string"
  description = "S3 bucket to store processed images"
}

variable "s3_bucket_prc_lifecycle" {
  type        = "string"
  description = "Enable or disable lifecycle rules on bucket with processed images. Values: true/false"
}

variable "s3_bucket_prc_lifecycle_expire" {
  type        = "string"
  description = "If lifecycle is enabled on on bucket with processed images, sets expiration time in days"
}

# AWS Lambda settings
variable "lambda_concurrency" {
  type        = "string"
  description = "AWS lambda concurrency"
}

variable "lambda_timeout" {
  type        = "string"
  description = "AWS lambda max execution time in seconds"
}

variable "lambda_memory" {
  type        = "string"
  description = "AWS lambda memory limit im MB"
}

variable "lambda_DEBUG" {
  type        = "string"
  description = "Mordo Debug mode. More logging, no S3 cache, CloudFront caching still applies."
}

variable "lambda_WATERMARK" {
  type        = "string"
  description = "Enable watermarking. true/false"
}

variable "lambda_WATERMARK_PATH" {
  type        = "string"
  description = "Watermark path if lambda_WATERMARK enabled."
}

variable "lambda_WATERMARK_SCALE" {
  type        = "string"
  description = "Watermark scale according to image size, if lambda_WATERMARK enabled."
}

variable "lambda_IMAGE_QUALITY" {
  type        = "string"
  description = "JPEG Image quality. 0-100"
}

variable "lambda_IMAGE_LOSSLESS" {
  type        = "string"
  description = "WEBP Image quality. true/false"
}

variable "lambda_IMAGE_COMPRESSION" {
  type        = "string"
  description = "PNG Image compression. 0-10"
}

variable "lambda_IMAGE_ENLARGE" {
  type        = "string"
  description = "Allow image enlarge during resize. true/false"
}

# AWS API Gateway settings
variable "apigw_throttling_rate_limit" {
  type        = "string"
  description = "AWS API Gateway throttling rate limit"
}

variable "apigw_throttling_burst_limit" {
  type        = "string"
  description = "AWS API Gateway throttling burst limit"
}

# CloudFront settings
variable "cf_price_class" {
  type        = "string"
  description = "CloudFront price class. Details: https://aws.amazon.com/cloudfront/pricing/"
}

variable "cf_min_ttl" {
  type        = "string"
  description = "The minimum amount of time (in seconds) that you want objects to stay in CloudFront caches before CloudFront queries your origin to see whether the object has been updated"
}

variable "cf_default_ttl" {
  type        = "string"
  description = "The default amount of time (in seconds) that an object is in a CloudFront cache before CloudFront forwards another request in the absence of an Cache-Control max-age or Expires header"
}
