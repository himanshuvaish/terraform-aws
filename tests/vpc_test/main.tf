provider "aws" {
  region = "us-west-2"
}

module "vpc" {
  source   = "../../modules/vpc"
  vpc_cidr = "10.0.0.0/16"
  vpc_name = "test-vpc"
}

output "vpc_id" {
  value = module.vpc.vpc_id
}
