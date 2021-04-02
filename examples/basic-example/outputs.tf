output vpc_id {
    value = module.basic_example.vpc_id
    description = "VPC ID"
}

output app_subnet_ids {
    value = module.basic_example.subnet_ids.app
    description = "App Subnet Ids"
}