resource "aws_cloudfront_distribution" "mordo" {
  count = "${length(var.s3_buckets)}"

  tags {
    Name = "${var.prefix}-mordo"
  }

  origin {
    domain_name = "${aws_api_gateway_rest_api.mordo.id}.execute-api.${data.aws_region.current.name}.amazonaws.com"

    origin_path = "${format("/%v/%v", 
      element(aws_api_gateway_stage.mordo.*.stage_name, count.index),
      aws_api_gateway_resource.mordo_process.path_part
    )}"

    origin_id = "${element(var.s3_buckets, count.index)}"

    custom_origin_config {
      http_port              = 80
      https_port             = 443
      origin_protocol_policy = "https-only"
      origin_ssl_protocols   = ["TLSv1.2"]
      origin_read_timeout    = 20
    }
  }

  enabled         = true
  is_ipv6_enabled = true
  comment         = "${var.prefix}-mordo"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "${element(var.s3_buckets, count.index)}"

    forwarded_values {
      query_string = true

      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = "${var.cf_min_ttl}"
    default_ttl            = "${var.cf_default_ttl}"
    max_ttl                = "${var.s3_bucket_prc_lifecycle_expire * 24 * 60 * 60}"
  }

  price_class = "${var.cf_price_class}"

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    cloudfront_default_certificate = true
  }
}

output "api-cloudfront-url" {
  value = "${formatlist("https://%s", aws_cloudfront_distribution.mordo.*.domain_name)}"
}
