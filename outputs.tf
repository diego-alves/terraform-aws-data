output vpc_id {
    value = data.aws_vpc.selected.id
    description = "ID da vpc"
}

output subnet_ids {
    value = data.aws_subnet_ids.selected.ids
    description = "Subnet ID"
}