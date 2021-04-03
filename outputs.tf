output "vpc_id" {
  value       = data.aws_vpc.selected.id
  description = "VPC ID"
}

output "subnet_ids" {
  value = tomap({
    for k, ids in data.aws_subnet_ids.selected : k => ids.ids
  })
  description = "Subnet ID"
}

output "account_id" {
  value       = data.aws_caller_identity.current.account_id
  description = "Account ID"
}

output "region" {
  value       = data.aws_region.current.name
  description = "Region"
}
