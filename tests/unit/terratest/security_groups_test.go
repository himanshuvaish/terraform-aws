package terratest

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupsModule(t *testing.T) {
	terraformOptions := &terraform.Options{
		TerraformDir: "../../../modules/security_groups",
		Vars: map[string]interface{}{
			"vpc_id":  "vpc-12345678",
			"sg_name": "test-sg",
		},
	}

	// Run terraform init and plan without applying the changes.
	planOutput := terraform.InitAndPlan(t, terraformOptions)

	// Assert that the plan includes an AWS security group resource.
	assert.True(t, strings.Contains(planOutput, "aws_security_group"),
		"Plan output should contain an aws_security_group resource definition")
}
