resource "aws_s3_bucket" "mordo_deploy" {
  bucket = "${var.prefix}-${var.s3_bucket_deploy}"
  acl    = "private"

  tags {
    Name = "${var.prefix}-${var.s3_bucket_deploy}"
  }

  lifecycle_rule {
    id      = "processed"
    enabled = "true"
    prefix  = "lambda/"

    expiration {
      days = "2"
    }
  }
}

resource "aws_s3_bucket" "mordo_processed" {
  bucket = "${var.prefix}-${var.s3_bucket_prc}"
  acl    = "private"

  tags {
    Name = "${var.prefix}-${var.s3_bucket_prc}"
  }

  lifecycle_rule {
    id      = "processed"
    enabled = "${var.s3_bucket_prc_lifecycle}"
    prefix  = "prc_"

    expiration {
      days = "${var.s3_bucket_prc_lifecycle_expire}"
    }
  }
}

output "AWS_S3_BUCKET_PROCESSED" {
  value = "${aws_s3_bucket.mordo_processed.bucket}"
}
