variable "vpc_logical_id" {
  description = "Tags for Vpc Identification"
  type = string
}

variable "subnet_logical_ids" {
  description = "Tags for Subnet Identification"
  type = map(list(string))
}