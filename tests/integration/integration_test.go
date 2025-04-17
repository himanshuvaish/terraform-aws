package terratest

import (
    "testing"

    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestIntegrationSetup(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "./integration_setup",
        Vars: map[string]interface{}{
            "public_subnet_id": "subnet-12345678", // Replace with your test subnet ID
        },
    }

    // Ensure resources are destroyed after testing.
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)

    // Retrieve outputs from Terraform.
    vpcID := terraform.Output(t, terraformOptions, "vpc_id")
    natGatewayID := terraform.Output(t, terraformOptions, "nat_gateway_id")
    routeTableID := terraform.Output(t, terraformOptions, "route_table_id")
    securityGroupID := terraform.Output(t, terraformOptions, "security_group_id")

    // Assert that outputs are not empty.
    assert.NotEmpty(t, vpcID, "VPC ID should not be empty")
    assert.NotEmpty(t, natGatewayID, "NAT Gateway ID should not be empty")
    assert.NotEmpty(t, routeTableID, "Route Table ID should not be empty")
    assert.NotEmpty(t, securityGroupID, "Security Group ID should not be empty")
}
