package terratest

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestRouteTablesModule(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../../modules/route_tables",
        Vars: map[string]interface{}{
            "vpc_id": "vpc-12345678",
            "nat_gateway_id": "nat-12345678",
            "public_route_table_name": "test-rt",
        },
    }
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    publicRouteTableID := terraform.Output(t, terraformOptions, "public_route_table_id")
    assert.NotEmpty(t, publicRouteTableID, "Public Route Table ID should not be empty")
}
