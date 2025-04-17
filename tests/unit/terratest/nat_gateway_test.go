package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestNatGatewayModule(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../../../modules/nat_gateway",
		Vars: map[string]interface{}{
			"public_subnet_id": "subnet-12345678",
			"nat_name":         "test-nat",
		},
	}

	// Run terraform init and plan instead of apply.
	planOutput := terraform.InitAndPlan(t, terraformOptions)

	// Assert that the plan includes an AWS Elastic IP resource.
	assert.True(t, strings.Contains(planOutput, "aws_eip"),
		"Plan output should contain an aws_eip resource definition")
}
