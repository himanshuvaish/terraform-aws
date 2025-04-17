variable "aws_region" {
  description = "AWS region"
  default     = "us-west-2"
}

variable "vpc_cidr" {
  description = "CIDR block for the VPC"
  default     = "10.0.0.0/16"
}

variable "vpc_name" {
  description = "Name of the VPC"
  default     = "netdevops-vpc"
}

# Provide the public subnet (this should be created or known beforehand)
variable "public_subnet_id" {
  description = "The public subnet id for NAT gateway"
  type        = string
}

variable "nat_name" {
  description = "Name of the NAT gateway"
  default     = "netdevops-nat"
}

variable "public_route_table_name" {
  description = "Name tag for the public route table"
  default     = "netdevops-public-rt"
}

variable "sg_name" {
  description = "Name tag for the security group"
  default     = "netdevops-sg"
}
