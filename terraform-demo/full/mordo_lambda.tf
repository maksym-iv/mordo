# data "archive_file" "mordo" {
#   type = "zip"

#   source_dir  = "../../pkg"
#   output_path = "../../pkg/mordo.zip"
# }

data "external" "mordo_zip" {
  program = ["sh", "./zip.sh", "../../pkg/mordo.zip", "../../pkg/"]
}

resource "aws_s3_bucket_object" "mordo_deploy" {
  bucket = "${aws_s3_bucket.mordo_deploy.bucket}"
  key    = "deploy/mordo.zip"
  source = "${data.external.mordo_zip.result.output_path}"
  acl    = "private"

  etag = "${data.external.mordo_zip.result.output_md5}"
}

resource "aws_lambda_function" "mordo" {
  function_name = "${var.prefix}-mordo"
  description   = "Mordo image service"

  tags {
    "Name" = "${var.prefix}-mordo"
  }

  s3_bucket = "${aws_s3_bucket_object.mordo_deploy.bucket}"
  s3_key    = "${aws_s3_bucket_object.mordo_deploy.key}"

  source_code_hash               = "${data.external.mordo_zip.result.output_base64sha256}"
  role                           = "${aws_iam_role.mordo_lambda.arn}"
  handler                        = "mordo"
  runtime                        = "go1.x"
  memory_size                    = "${var.lambda_memory}"
  timeout                        = "${var.lambda_timeout}"
  reserved_concurrent_executions = "${var.lambda_concurrency}"
  publish                        = false

  environment {
    variables {
      DEBUG                   = "${var.lambda_DEBUG}"
      PLATFORM                = "aws"
      MAXMEM                  = "${var.lambda_memory - 20}"
      CACHE_MAXAGE            = "${var.s3_bucket_prc_lifecycle_expire * 24 * 60 * 60}"
      AWS_S3_BUCKET_PROCESSED = "${aws_s3_bucket.mordo_processed.bucket}"
      WATERMARK               = "${var.lambda_WATERMARK}"
      WATERMARK_PATH          = "${var.lambda_WATERMARK_PATH}"
      WATERMARK_SCALE         = "${var.lambda_WATERMARK_SCALE}"
      IMAGE_QUALITY           = "${var.lambda_IMAGE_QUALITY}"
      IMAGE_LOSSLESS          = "${var.lambda_IMAGE_LOSSLESS}"
      IMAGE_COMPRESSION       = "${var.lambda_IMAGE_COMPRESSION}"
      IMAGE_ENLARGE           = "${var.lambda_IMAGE_ENLARGE}"
    }
  }
}

resource "aws_lambda_alias" "mordo_deploy" {
  name             = "deployed"
  description      = "${var.prefix} stage"
  function_name    = "${aws_lambda_function.mordo.arn}"
  function_version = "${aws_lambda_function.mordo.version}"
}

resource "aws_lambda_permission" "apigw_lambda" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.mordo.function_name}"

  qualifier = "${aws_lambda_alias.mordo_deploy.name}"
  principal = "apigateway.amazonaws.com"

  # More: http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html
  source_arn = "arn:aws:execute-api:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.mordo.id}/*/${aws_api_gateway_method.mordo_process_image.http_method}/${aws_api_gateway_resource.mordo_process.path_part}/${aws_api_gateway_resource.mordo_process_image.path_part}"
}
