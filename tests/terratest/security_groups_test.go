package terratest

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestSecurityGroupsModule(t *testing.T) {
    terraformOptions := &terraform.Options{
        TerraformDir: "../../modules/security_groups",
        Vars: map[string]interface{}{
            "vpc_id": "vpc-12345678",
            "sg_name": "test-sg",
        },
    }
    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
    securityGroupID := terraform.Output(t, terraformOptions, "security_group_id")
    assert.NotEmpty(t, securityGroupID, "Security Group ID should not be empty")
}
