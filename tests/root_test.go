package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestRootExample(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "..",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": "us-east-1",
		},

		Vars: map[string]interface{}{
			"vpc_logical_id": "MainVPC",
			"subnet_logical_ids": map[string][]string{
				"app": {"AppSubNet1", "AppSubNet2"},
				"pub": {"PublicSubNet1", "PublicSubNet2"},
				"db":  {"DBSubNet1", "DBSubNet2"},
			},
		},

		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	subnetIds := terraform.OutputMap(t, terraformOptions, "subnet_ids")

	assert.Equal(t, "vpc-", vpcId[:4])
	for k, v := range subnetIds {
		assert.Contains(t, "app pub db", k)
		for _, el := range strings.Split(v[1:8], " ") {
			assert.Equal(t, "subnet-", strings.ReplaceAll(el, "[", "")[:7])
		}
	}
}
