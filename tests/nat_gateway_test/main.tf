provider "aws" {
  region = "us-west-2"
}

module "nat_gateway" {
  source           = "../../modules/nat_gateway"
  public_subnet_id = "subnet-12345678"  # Dummy subnet id for test purposes
  nat_name         = "test-nat"
}

output "nat_gateway_id" {
  value = module.nat_gateway.nat_gateway_id
}
