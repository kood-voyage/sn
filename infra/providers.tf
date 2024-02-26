terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    # here we are storing our Terraform infrastructure state file in an encrypted S3 bucket
    bucket = "terraform-state-koodvoyage-sn"
    key    = "global/s3/terraform.tfstate"
    region = "us-east-1"
    #dynamodb_table = "terraform-koodvoyage-sn-state-locks"
    encrypt = true
  }


}
provider "aws" {
  region = "us-east-1"
}