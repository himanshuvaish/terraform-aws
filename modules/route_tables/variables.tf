variable "vpc_id" {
  description = "VPC ID"
  type        = string
}

variable "nat_gateway_id" {
  description = "NAT gateway ID"
  type        = string
}

variable "public_route_table_name" {
  description = "Name for the public route table"
  type        = string
}
