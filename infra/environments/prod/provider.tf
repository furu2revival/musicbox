provider "aws" {
  region = var.region

  default_tags {
    tags = {
      "Terraform"   = "true"
      "Project"     = "musicbox"
      "Environment" = var.env
    }
  }
}

provider "aws" {
  region = "us-east-1"
  alias  = "virginia"
}

provider "awscc" {
  region = var.region
}
