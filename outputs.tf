output "vpc_id" {
  value = module.vpc.vpc_id
}

output "nat_gateway_id" {
  value = module.nat_gateway.nat_gateway_id
}

output "public_route_table_id" {
  value = module.route_tables.public_route_table_id
}

output "security_group_id" {
  value = module.security_groups.security_group_id
}
