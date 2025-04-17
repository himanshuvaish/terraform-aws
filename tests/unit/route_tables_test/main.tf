provider "aws" {
  region = "us-west-2"
}

# For route table testing, we first instantiate the VPC and NAT gateway modules to provide dependencies.
module "vpc" {
  source   = "../../modules/vpc"
  vpc_cidr = "10.0.0.0/16"
  vpc_name = "test-vpc"
}

module "nat_gateway" {
  source           = "../../modules/nat_gateway"
  public_subnet_id = "subnet-12345678"  # Dummy subnet id for test purposes
  nat_name         = "test-nat"
}

module "route_tables" {
  source                  = "../../modules/route_tables"
  vpc_id                  = module.vpc.vpc_id
  nat_gateway_id          = module.nat_gateway.nat_gateway_id
  public_route_table_name = "test-public-rt"
}

output "public_route_table_id" {
  value = module.route_tables.public_route_table_id
}
