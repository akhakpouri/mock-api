output "app_url" {
  description = "app url"
  value       = aws_alb.application_lobad_balancer.dns_name
}
