terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

terraform {
  backend "s3" {
    # here we are storing our terraform infrastructure state file in an encrypted s3 bucket that we declared in main.tf
    bucket         = "terraform-state-koodvoyage-sn"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"

    dynamodb_table = "terraform-koodvoyage-sn-state-locks"
    encrypt        = true
  }
}

provider "aws" {
  region                   = "us-east-1"
  shared_credentials_files = ["~/.aws/credentials"]
  profile                  = "vscode"
}