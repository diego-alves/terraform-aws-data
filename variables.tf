variable "vpc_tags" {
  description = "Tags for Vpc Identification"
  type = map(string)
}

variable "subnet_tags" {
  description = "Tags for Subnet Identification"
  type = map(string)
}