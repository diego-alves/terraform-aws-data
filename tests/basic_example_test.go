package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformBasicExample(t *testing.T) {
	t.Parallel()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic-example",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": "us-east-1",
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	appSubnetIds := terraform.OutputList(t, terraformOptions, "app_subnet_ids")

	assert.Equal(t, "vpc-", vpcId[:4])
	assert.Equal(t, 2, len(appSubnetIds))
	for _, element := range appSubnetIds {
		assert.Equal(t, "subnet-", element[:7])
	}

}
