data "aws_vpc" "selected" {
  tags = var.vpc_tags
}

data "aws_subnet_ids" "selected" {
  vpc_id = data.aws_vpc.selected.id
  tags = {
    "aws:cloudformation:logical-id" : ""
  }
}