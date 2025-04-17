package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRouteTablesModule(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../../modules/route_tables",
		Vars: map[string]interface{}{
			"vpc_id":                  "vpc-12345678",
			"nat_gateway_id":          "nat-12345678",
			"public_route_table_name": "test-rt",
		},
	}

	// Execute terraform init and plan in plan-only mode.
	planOutput := terraform.InitAndPlan(t, terraformOptions)

	// Assert that the plan includes an AWS route table resource (for example, check for aws_route_table)
	assert.True(t, strings.Contains(planOutput, "aws_route_table"),
		"Plan output should contain an aws_route_table resource definition")
}
