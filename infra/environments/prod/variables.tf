variable "env" {
  description = "環境名"
  type        = string
  default     = "prod"
}

variable "region" {
  description = "リージョン"
  type        = string
  default     = "ap-northeast-1"
}

variable "server-version" {
  description = "musicbox/api-server の docker image tag"
  type        = string
  default     = "latest"
}
