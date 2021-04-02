module "basic_example" {
  source = "../.."
  vpc_logical_id = "MainVPC"
	subnet_logical_ids = {
		app: ["AppSubNet1", "AppSubNet2"]
		pub: ["PublicSubNet1", "PublicSubNet2"]
		db:  ["DBSubNet1", "DBSubNet2"]
	}
}
