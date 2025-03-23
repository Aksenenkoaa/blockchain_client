variable "region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "aws_region" {
  description = "The AWS region to deploy resources"
  type        = string
  default     = "us-west-2"
}

variable "app_name" {
  description = "The name of the application"
  type        = string
  default     = "blockchain-client"
}

variable "app_image" {
  description = "The Docker image for the application"
  type        = string
  default     = "blockchain-client:latest"
}

variable "app_port" {
  description = "The port the application listens on"
  type        = number
  default     = 8080
}

variable "ecs_task_cpu" {
  description = "The CPU units for the ECS task"
  type        = number
  default     = 256
}

variable "ecs_task_memory" {
  description = "The memory for the ECS task"
  type        = number
  default     = 512
}

variable "vpc_id" {
  description = "The VPC ID where resources will be deployed"
  type        = string
}

variable "public_subnets" {
  description = "List of public subnets for the ECS service"
  type        = list(string)
}

variable "security_group_id" {
  description = "Security group ID for the ECS service"
  type        = string
}