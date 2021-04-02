output vpc_id {
    value = module.basic_example.vpc_id
    description = "ID da vpc"
}

output subnet_id {
    value = module.basic_example.subnet_id
    description = "Subnet ID"
}