data "aws_vpc" "selected" {
  tags = {
    "aws:cloudformation:logical-id" : var.vpc_logical_id
  }
}

data "aws_subnet_ids" "selected" {
  for_each = var.subnet_logical_ids
  vpc_id = data.aws_vpc.selected.id
  filter {
    name   = "tag:aws:cloudformation:logical-id"
    values = each.value
  }
}