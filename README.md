[![Maintained by Diego Alves](https://img.shields.io/badge/maintained%20by-diego--alves-blue)](https://github.com/diego-alves) [![Module Version](https://img.shields.io/github/v/tag/diego-alves/terraform-aws-data?label=version&sort=semver)](https://registry.terraform.io/modules/diego-alves/data/aws/latest)

# AWS Data Terraform Module

This repo contains a [Terraform](https://terraform.io) [Module](https://www.terraform.io/docs/language/modules/index.html) for retrieving some data from the current account the automation is being applied.

## How to use this Module

For the basic usage you can include as follows:

```hcl
module "data" {
  source  = "diego-alves/data/aws"
  version = "0.0.5"
}

```

Now you can reference the subnet or vcp as `module.data.subnet_ids.app` or `module.data.vpc_id`.

The basic usage assumes you have a VPC and Subnets with Cloudformation Tag `aws:cloudformation:logical-id`, and the values for these tags are:

- `MainVpc` for the VPC as `vpc_id` in the output.
- `AppSubNet1` and `AppSubNet1` for the application subnets as `subnet_ids.app` in the output.
- `PublicSubNet1` and `PublicSubNet1` for the public subnets as `subnet_ids.pub` in the output.
- `DBSubNet1` and `DBSubNet1` for the data base subnets as `subnet_ids.dat` in the output.

If your tag is different, you can specify the variables as follows:

```hcl
module "data" {
  source  = "diego-alves/data/aws"
  version = "0.0.5"

  vpc_logical_id = "MainVPC"
  subnet_logical_ids = {
    app : ["AppSubNet1", "AppSubNet2"]
    pub : ["PublicSubNet1", "PublicSubNet2"]
    dat : ["DBSubNet1", "DBSubNet2"]
  }
}
```
