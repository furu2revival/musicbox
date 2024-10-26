variable "env" {
  type = string
}

variable "region" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "private_subnet_ids" {
  type = list(string)
}

variable "server-version" {
  description = "musicbox/api-server の docker image tag"
  type        = string
  default     = "latest"
}

variable "lb_target_group_arn" {
  type = string
}

