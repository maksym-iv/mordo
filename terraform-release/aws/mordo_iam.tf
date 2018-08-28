data "aws_iam_policy_document" "mordo_lambda" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = [
      "sts:AssumeRole",
    ]
  }
}

resource "aws_iam_role" "mordo_lambda" {
  name               = "${var.prefix}-mordo-lambda"
  assume_role_policy = "${data.aws_iam_policy_document.mordo_lambda.json}"
}

resource "aws_iam_role_policy_attachment" "mordo_lambda" {
  role       = "${aws_iam_role.mordo_lambda.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

data "aws_iam_policy_document" "mordo_lambda_s3" {
  statement {
    sid    = 1
    effect = "Allow"

    actions = [
      "s3:PutObjectTagging",
      "s3:PutObject",
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObjectAcl",
    ]

    resources = [
      "${aws_s3_bucket.mordo_processed.arn}",
      "${aws_s3_bucket.mordo_processed.arn}/*",
    ]
  }

  statement {
    sid    = 2
    effect = "Allow"

    actions = [
      "s3:GetObject",
      "s3:ListBucket",
      "s3:PutObjectAcl",
    ]

    resources = [
      "${formatlist("arn:aws:s3:::%s", var.s3_buckets)}",
      "${formatlist("arn:aws:s3:::%s/*", var.s3_buckets)}",
    ]
  }

  statement {
    sid    = 3
    effect = "Allow"

    actions = [
      "s3:GetBucketLocation",
      "s3:ListAllMyBuckets",
    ]

    resources = [
      "*",
    ]
  }
}

resource "aws_iam_role_policy" "mordo_lambda_s3" {
  role   = "${aws_iam_role.mordo_lambda.name}"
  name   = "mordo_lambda_s3"
  policy = "${data.aws_iam_policy_document.mordo_lambda_s3.json}"
}
