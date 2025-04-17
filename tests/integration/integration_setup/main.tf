module "vpc" {
  source = "../../modules/vpc"
  vpc_cidr = "10.0.0.0/16"
  vpc_name = "integration-vpc"
}

module "nat_gateway" {
  source = "../../modules/nat_gateway"
  public_subnet_id = var.public_subnet_id
  nat_name = "integration-nat"
}

module "route_tables" {
  source = "../../modules/route_tables"
  vpc_id = module.vpc.vpc_id
  nat_gateway_id = module.nat_gateway.nat_gateway_id
  public_route_table_name = "integration-rt"
}

module "security_groups" {
  source = "../../modules/security_groups"
  vpc_id = module.vpc.vpc_id
  sg_name = "integration-sg"
}
