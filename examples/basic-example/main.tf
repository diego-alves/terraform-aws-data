module "basic_example" {
  source = "../.."
  vpc_tags = {
    "aws:cloudformation:logical-id" : "MainVPC"
  }
  subnet_tags = {
    "aws:cloudformation:logical-id" : "AppSubNet2"
  }
}
