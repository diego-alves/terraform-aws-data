variable "vpc_logical_id" {
  description = "Tags for Vpc Identification"
  type        = string
  default     = "MainVPC"
}

variable "subnet_logical_ids" {
  description = "Tags for Subnet Identification"
  type        = map(list(string))
  default = {
    app : ["AppSubNet1", "AppSubNet2"]
    pub : ["PublicSubNet1", "PublicSubNet2"]
    dat : ["DBSubNet1", "DBSubNet2"]
  }
}
