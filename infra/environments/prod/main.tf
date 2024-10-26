module "vpc" {
  source = "../../modules/vpc"

  env = var.env
}

module "ec2" {
  source = "../../modules/ec2"

  env              = var.env
  vpc_id           = module.vpc.id
  public_subnet_id = module.vpc.public_subnet_ids[0]
}

module "rds" {
  source = "../../modules/rds"

  env                              = var.env
  vpc_id                           = module.vpc.id
  private_subnet_ids               = module.vpc.private_subnet_ids
  bastion_server_security_group_id = module.ec2.security_group_id
  instance_type                    = "db.m7g.large"
}

module "elb" {
    source = "../../modules/elb"

    env               = var.env
    vpc_id            = module.vpc.id
    public_subnet_ids = module.vpc.public_subnet_ids
}

module "ecs" {
  source = "../../modules/ecs"

  env                   = var.env
  region                = var.region
  vpc_id                = module.vpc.id
  private_subnet_ids    = module.vpc.private_subnet_ids
  lb_target_group_arn   = module.elb.lb_target_group_arn
}
