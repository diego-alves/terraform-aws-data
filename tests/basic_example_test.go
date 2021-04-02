package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformBasicExample(t *testing.T) {
	t.Parallel()
	expectedRegion := "us-east-1"

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/basic-example",
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": expectedRegion,
		},
		NoColor: true,
	})

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	vpcId := terraform.Output(t, terraformOptions, "vpc_id")
	appSubnetIds := terraform.OutputList(t, terraformOptions, "app_subnet_ids")
	pubSubnetIds := terraform.OutputList(t, terraformOptions, "pub_subnet_ids")
	dbSubnetIds := terraform.OutputList(t, terraformOptions, "db_subnet_ids")
	accountId := terraform.Output(t, terraformOptions, "account_id")
	region := terraform.Output(t, terraformOptions, "region")

	assert.Equal(t, expectedRegion, region)
	assert.Regexp(t, "^\\d{12}$", accountId)
	assert.Regexp(t, "^vpc-[0-9a-f]{17}$", vpcId)
	assert.Equal(t, 2, len(appSubnetIds))
	assert.Equal(t, 2, len(pubSubnetIds))
	assert.Equal(t, 2, len(dbSubnetIds))

	for _, element := range appSubnetIds {
		assert.Regexp(t, "^subnet-[0-9a-f]{17}$", element)
	}

	for _, element := range pubSubnetIds {
		assert.Regexp(t, "^subnet-[0-9a-f]{17}$", element)
	}

	for _, element := range dbSubnetIds {
		assert.Regexp(t, "^subnet-[0-9a-f]{17}$", element)
	}

}
