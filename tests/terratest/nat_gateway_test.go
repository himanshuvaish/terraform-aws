package terratest

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestNatGatewayModule(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../../modules/nat_gateway",
        Vars: map[string]interface{}{
            "public_subnet_id": "subnet-12345678",
            "nat_name": "test-nat",
        },
    }
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    natGatewayID := terraform.Output(t, terraformOptions, "nat_gateway_id")
    assert.NotEmpty(t, natGatewayID, "NAT Gateway ID should not be empty")
}
