terraform {
    required_providers {
      aws = {
        source = "hashicorp/aws"
        version = "~> 5.0"
      }
    }
}

provider "aws" {
  region = var.aws_region
}

variable "aws_region" {
  description = "AWS region to deploy resources"
  type = string
  default = "us-east-1"
}

variable "bucket_name" {
  description = "Name of the S3 bucket for workflow template storage"
  type = string
  default = "seryn-workflow-templates"
}

resource "aws_s3_bucket" "workflow_templates" {
  bucket = var.bucket_name
  tags = {
    Project = "Seryn"
    Purpose = "Workflow template storage"
    Environment = "simulation"
    ManagedBy = "Terraform"
  }
}

resource "aws_s3_bucket_versioning" "workflow_templates" {
  bucket = aws_s3_bucket.workflow_templates.id

  versioning_configuration {
    status = "Enabled"
  }
}

resource "aws_s3_bucket_server_side_encryption_configuration" "workflow_templates" {
    bucket = aws_s3_bucket.workflow_templates.id

    rule {
        apply_server_side_encryption_by_default {
          sse_algorithm = "AES256"
        }
    } 
}

resource "aws_s3_bucket_public_access_block" "workflow_templates" {
  bucket = aws_s3_bucket.workflow_templates.id

  block_public_acls = true
  block_public_policy = true
  ignore_public_acls = true
  restrict_public_buckets = true
}

output "bucket_name" {
    description = "Name of the created s3 bucket"
    value = aws_s3_bucket.workflow_templates.bucket
}

output "bucket_arn" {
    description = "ARN of the created s3 bucket"
    value = aws_s3_bucket.workflow_templates.arn
}