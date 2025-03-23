# Blockchain Client Deployment on AWS ECS Fargate

This project deploys a simple blockchain client application to AWS ECS Fargate using Terraform. The application exposes an API to interact with the Polygon blockchain.

## Prerequisites

Before deploying the application, ensure you have the following:

1. **Terraform**: Install Terraform from [here](https://www.terraform.io/downloads.html).
2. **AWS CLI**: Install and configure the AWS CLI with your credentials:

Run in docker
```
Docker run -p 8080:8080 blockchain-client
```

## Deployment Steps

Clone the Repository:
bash
```
git clone https://github.com/your-repo/blockchain-client.git
cd blockchain-client/terraform
```

Initialize Terraform:
bash
```
terraform init
```

Review the Execution Plan:
bash
```
terraform plan
```

Deploy the Infrastructure:
bash
```
terraform apply
```

Access the Application:
After deployment, Terraform will output the ALB DNS name. Use this to access the application:
bash
```
curl http://<alb_dns_name>/blockNumber
```
