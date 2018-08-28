resource "aws_api_gateway_rest_api" "mordo" {
  name        = "${var.prefix}-mordo"
  description = "This is API GW for ${var.prefix} mordo service"

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  binary_media_types = [
    "*/*",
    "image/*",
    "application/octet-stream",
  ]
}

// resource "aws_api_gateway_request_validator" "mordo" {
//   name                        = "mordo"
//   rest_api_id                 = "${aws_api_gateway_rest_api.mordo.id}"
//   validate_request_body       = false
//   validate_request_parameters = true
// }

resource "aws_api_gateway_resource" "mordo_process" {
  rest_api_id = "${aws_api_gateway_rest_api.mordo.id}"
  parent_id   = "${aws_api_gateway_rest_api.mordo.root_resource_id}"
  path_part   = "process"
}

resource "aws_api_gateway_resource" "mordo_process_image" {
  rest_api_id = "${aws_api_gateway_rest_api.mordo.id}"
  parent_id   = "${aws_api_gateway_resource.mordo_process.id}"
  path_part   = "{image+}"
}

resource "aws_api_gateway_method" "mordo_process_image" {
  rest_api_id   = "${aws_api_gateway_rest_api.mordo.id}"
  resource_id   = "${aws_api_gateway_resource.mordo_process_image.id}"
  http_method   = "GET"
  authorization = "NONE"

  // request_parameters {
  //   method.request.querystring.format = true
  //   method.request.querystring.ip     = true
  // }

  // request_validator_id = "${aws_api_gateway_request_validator.mordo.id}"
}

resource "aws_api_gateway_integration" "mordo" {
  rest_api_id             = "${aws_api_gateway_rest_api.mordo.id}"
  resource_id             = "${aws_api_gateway_resource.mordo_process_image.id}"
  http_method             = "${aws_api_gateway_method.mordo_process_image.http_method}"
  type                    = "AWS_PROXY"
  integration_http_method = "POST"

  uri = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_alias.mordo_deploy.arn}/invocations"
}

data "aws_s3_bucket" "mordo_images" {
  count  = "${length(var.s3_buckets)}"
  bucket = "${element(var.s3_buckets, count.index)}"
}

resource "aws_api_gateway_deployment" "mordo" {
  count             = "${length(var.s3_buckets)}"
  stage_name        = "deployment-${element(var.s3_buckets, count.index)}"
  rest_api_id       = "${aws_api_gateway_rest_api.mordo.id}"
  stage_description = "Deployed at ${timestamp()}"

  variables {
    s3_bucket        = "${element(var.s3_buckets, count.index)}"
    s3_bucket_region = "${element(data.aws_s3_bucket.mordo_images.*.region, count.index)}"
  }

  depends_on = [
    "aws_api_gateway_integration.mordo",
  ]
}

resource "aws_api_gateway_method_settings" "mordo_deployment" {
  count       = "${length(var.s3_buckets)}"
  rest_api_id = "${aws_api_gateway_rest_api.mordo.id}"
  stage_name  = "${element(aws_api_gateway_deployment.mordo.*.stage_name, count.index)}"
  method_path = "*/*"

  settings {
    throttling_rate_limit  = "1"
    throttling_burst_limit = "1"
    caching_enabled        = "false"
  }
}

resource "aws_api_gateway_stage" "mordo" {
  count         = "${length(var.s3_buckets)}"
  stage_name    = "${element(var.s3_buckets, count.index)}"
  rest_api_id   = "${aws_api_gateway_rest_api.mordo.id}"
  deployment_id = "${aws_api_gateway_deployment.mordo.id}"

  variables {
    s3_bucket        = "${element(var.s3_buckets, count.index)}"
    s3_bucket_region = "${element(data.aws_s3_bucket.mordo_images.*.region, count.index)}"
  }
}

resource "aws_api_gateway_method_settings" "mordo" {
  count       = "${length(var.s3_buckets)}"
  rest_api_id = "${aws_api_gateway_rest_api.mordo.id}"
  stage_name  = "${element(aws_api_gateway_stage.mordo.*.stage_name, count.index)}"
  method_path = "*/*"

  settings {
    throttling_rate_limit  = "${var.apigw_throttling_rate_limit}"
    throttling_burst_limit = "${var.apigw_throttling_burst_limit}"
  }
}

output "api-gw-url" {
  value = "${aws_api_gateway_stage.mordo.*.invoke_url}"
}
