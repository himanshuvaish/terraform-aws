package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestVpcModule(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../../modules/vpc",
		Vars: map[string]interface{}{
			"vpc_cidr": "10.0.0.0/16",
			"vpc_name": "test-vpc",
		},
	}

	// Run terraform init and plan instead of applying real changes.
	planOutput := terraform.InitAndPlan(t, terraformOptions)

	// Assert that the plan includes an AWS VPC resource.
	assert.True(t, strings.Contains(planOutput, "aws_vpc"),
		"Plan output should contain an aws_vpc resource definition")
}
