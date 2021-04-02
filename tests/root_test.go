package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRootExample(t *testing.T) {
	t.Parallel()
	expectedRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "..",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": expectedRegion,
		},

		Vars: map[string]interface{}{
			"vpc_logical_id": "MainVPC",
			"subnet_logical_ids": map[string][]string{
				"app": {"AppSubNet1", "AppSubNet2"},
				"pub": {"PublicSubNet1", "PublicSubNet2"},
				"dat": {"DBSubNet1", "DBSubNet2"},
			},
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	subnetIds := terraform.OutputMap(t, terraformOptions, "subnet_ids")
	accountId := terraform.Output(t, terraformOptions, "account_id")
	region := terraform.Output(t, terraformOptions, "region")

	assert.Equal(t, expectedRegion, region)
	assert.Regexp(t, "^\\d{12}$", accountId)
	assert.Regexp(t, "^vpc-[0-9a-f]{17}$", vpcId)
	assert.Regexp(t, "^\\[subnet-[0-9a-f]{17} subnet-[0-9a-f]{17}\\]$", subnetIds["app"])
	assert.Regexp(t, "^\\[subnet-[0-9a-f]{17} subnet-[0-9a-f]{17}\\]$", subnetIds["pub"])
	assert.Regexp(t, "^\\[subnet-[0-9a-f]{17} subnet-[0-9a-f]{17}\\]$", subnetIds["dat"])

}
