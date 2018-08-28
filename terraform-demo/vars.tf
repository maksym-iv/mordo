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

variable "s3_bucket_prc_lifecycle_expire" {
  type        = "string"
  description = "Enable or disable lifecycle rules on bucket with processed images. Values: true/false"
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
