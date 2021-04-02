module "basic_example" {
  source = "../.."
  vpc_tags = {
    "aws:cloudformation:logical-id" : ""
  }
}
