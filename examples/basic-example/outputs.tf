output vpc_id {
    value = module.basic_example.vpc_id
    description = "VPC ID"
}

output app_subnet_ids {
    value = module.basic_example.subnet_ids.app
    description = "App Subnet Ids"
}

output pub_subnet_ids {
    value = module.basic_example.subnet_ids.pub
    description = "Public Subnet Ids"
}

output db_subnet_ids {
    value = module.basic_example.subnet_ids.db
    description = "Data Subnet Ids"
}