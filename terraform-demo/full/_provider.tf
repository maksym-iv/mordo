provider "aws" {
  // profile = "terraform"

  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"

  region = "${var.aws_region}"
}

provider "aws" {
  # Used for serverless certs. (API GW, CloudFron require ecrt from us-east-1)
  alias = "us-east-1"

  // profile = "terraform"

  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "us-east-1"
}
