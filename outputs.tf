output vpc_id {
    value = data.aws_vpc.selected.id
    description = "ID da vpc"
}

output subnet_ids {
    value = tomap({
        for k, ids in data.aws_subnet_ids.selected: k => ids.ids
    })
    description = "Subnet ID"
}