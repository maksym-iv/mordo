resource "aws_cloudfront_distribution" "mordo_demo" {
  tags {
    Name = "demo-mordo"
  }

  aliases = ["demo.mordo.io"]

  origin {
    domain_name = "cmmvvx6w13.execute-api.us-west-1.amazonaws.com"
    origin_path = "/dev-mordo-images/process"
    origin_id   = "demo-mordo"

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
  comment         = "demo-mordo"

  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "demo-mordo"

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
    acm_certificate_arn      = "arn:aws:acm:us-east-1:962546435606:certificate/b022e814-2503-44aa-881d-6cf552ce4d7c"
    ssl_support_method       = "sni-only"
    minimum_protocol_version = "TLSv1.1_2016"

    # minimum_protocol_version = "TLSv1.2_2018"
  }
}

resource "aws_route53_record" "mordo_demo" {
  zone_id = "Z1OP2PQEZH7O0H"
  name    = "demo"
  type    = "A"

  alias {
    name                   = "${aws_cloudfront_distribution.mordo_demo.domain_name}"
    zone_id                = "${aws_cloudfront_distribution.mordo_demo.hosted_zone_id}"
    evaluate_target_health = false
  }
}

output "api-cf-url" {
  value = "${formatlist("https://%s", aws_cloudfront_distribution.mordo_demo.*.domain_name)}"
}

output "api-domain" {
  value = "https://${aws_route53_record.mordo_demo.fqdn}"
}

output "image-sample" {
  value = "https://${aws_route53_record.mordo_demo.fqdn}/origin_2k.jpg?width=1000"
}
