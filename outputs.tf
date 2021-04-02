output vpc_id {
    value = data.aws_vpc.selected.id
    description = "ID da vpc"
}

output subnet_id {
    value = data.aws_subnet.selected.id
    description = "Subnet ID"
}