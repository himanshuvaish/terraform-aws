provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source   = "./modules/vpc"
  vpc_cidr = var.vpc_cidr
  vpc_name = var.vpc_name
}

module "nat_gateway" {
  source           = "./modules/nat_gateway"
  public_subnet_id = var.public_subnet_id
  nat_name         = var.nat_name
}

module "route_tables" {
  source                  = "./modules/route_tables"
  vpc_id                  = module.vpc.vpc_id
  nat_gateway_id          = module.nat_gateway.nat_gateway_id
  public_route_table_name = var.public_route_table_name
}

module "security_groups" {
  source  = "./modules/security_groups"
  vpc_id  = module.vpc.vpc_id
  sg_name = var.sg_name
}
