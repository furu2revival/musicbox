variable "env" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "private_subnet_ids" {
  type = list(string)
}

variable "bastion_server_security_group_id" {
  description = "踏み台サーバのセキュリティグループのID"
  type        = string
}

variable "instance_type" {
  type    = string
  default = "db.t2.micro"
}
