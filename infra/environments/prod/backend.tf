terraform {
  backend "s3" {
    bucket         = "musicbox-tf-state-prod"
    key            = "terraform.tfstate"
    region         = "ap-northeast-1"
    dynamodb_table = "musicbox-tf-state-lock-prod"
  }
}
