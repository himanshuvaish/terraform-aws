provider "aws" {
  region = "us-west-2"
}

# For security groups, we need a VPC; instantiate the VPC module first.
module "vpc" {
  source   = "../../modules/vpc"
  vpc_cidr = "10.0.0.0/16"
  vpc_name = "test-vpc"
}

module "security_groups" {
  source  = "../../modules/security_groups"
  vpc_id  = module.vpc.vpc_id
  sg_name = "test-sg"
}

output "security_group_id" {
  value = module.security_groups.security_group_id
}
