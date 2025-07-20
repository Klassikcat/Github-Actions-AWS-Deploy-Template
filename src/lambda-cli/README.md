# PoorServerless

A modern CLI tool for generating serverless functions with AWS CDK deployment infrastructure.

## Features

- 🚀 Generate serverless functions for multiple runtimes (Python, Node.js, Go)
- ☁️ AWS CDK v2 infrastructure as code (TypeScript alternative in Go)
- 📦 Complete CloudFormation stack with API Gateway + Lambda
- 🛠️ Pre-configured with production-ready templates
- 📋 One-command deployment scripts

## Why CDK over SAM?

- **More flexible**: Programmatic infrastructure definition
- **Better abstractions**: Higher-level constructs for complex scenarios
- **Multi-service support**: Easy integration with other AWS services
- **Type safety**: Compile-time checks for infrastructure code
- **Extensible**: Custom constructs and reusable patterns

## Installation

1. Build the CLI tool:
```bash
cd src/PoorServerless
go build -o poorserverless .
```

2. (Optional) Add to PATH:
```bash
# Add to your shell profile (.bashrc, .zshrc, etc.)
export PATH=$PATH:/path/to/PoorServerless
```

## Prerequisites

- **AWS CLI** configured with appropriate credentials
- **AWS CDK CLI** installed: `npm install -g aws-cdk`
- **Go 1.21+** for building the tool and Go functions

## Usage

### Create a new serverless function

```bash
# Create a Python function with CDK infrastructure
./poorserverless create-function --name my-python-func --runtime python --output ./functions

# Create a Node.js function with CDK infrastructure
./poorserverless create-function --name my-node-func --runtime nodejs --output ./functions

# Create a Go function with CDK infrastructure
./poorserverless create-function --name my-go-func --runtime go --output ./functions
```

### Command Options

- `--name, -n`: Function name (required)
- `--runtime, -r`: Runtime environment (python, nodejs, go) [default: python]
- `--output, -o`: Output directory [default: current directory]

## Generated Structure

```
my-function/
├── app.py|index.js|main.go    # Your Lambda function code
├── requirements.txt|package.json|go.mod  # Dependencies
├── infrastructure/            # CDK Infrastructure
│   ├── main.go               # CDK stack definition
│   ├── go.mod                # CDK dependencies
│   └── cdk.json              # CDK configuration
├── deploy.sh                 # One-command deployment
└── README.md                 # Function-specific documentation
```

## Infrastructure Components

Each generated function includes:

- **AWS Lambda Function** with your code
- **API Gateway REST API** with CORS enabled
- **IAM Roles and Policies** with least privilege
- **CloudWatch Logs** for monitoring
- **CloudFormation Outputs** for easy access to endpoints

## Deployment

After generating a function:

```bash
# Navigate to your function directory
cd my-function

# Deploy everything with one command
./deploy.sh
```

Or manually:

```bash
cd infrastructure

# Install dependencies
go mod tidy

# Bootstrap CDK (once per account/region)
cdk bootstrap

# Deploy the stack
cdk deploy
```

## Development Workflow

1. **Generate function**: `./poorserverless create-function --name my-api --runtime python`
2. **Develop locally**: Modify your function code and test locally
3. **Update infrastructure**: Edit `infrastructure/main.go` if needed
4. **Deploy**: Run `./deploy.sh`
5. **Test**: Use the API Gateway URL from the output

## Advanced Features

### Custom Infrastructure

Modify `infrastructure/main.go` to add:

- **Environment variables**
- **VPC configuration**
- **Database connections**
- **SQS/SNS integration**
- **Custom IAM policies**
- **CloudWatch alarms**

### Multiple Environments

```bash
# Deploy to different environments
cd infrastructure
cdk deploy --context env=dev
cdk deploy --context env=prod
```

## Examples

```bash
# E-commerce API
./poorserverless create-function --name product-catalog --runtime python

# Real-time notifications
./poorserverless create-function --name notification-service --runtime nodejs

# High-performance data processing
./poorserverless create-function --name data-processor --runtime go
```

## Cleanup

To remove all resources:

```bash
cd infrastructure
cdk destroy
```

## Comparison with SAM

| Feature | SAM | PoorServerless + CDK |
|---------|-----|---------------------|
| Learning Curve | Easy | Moderate |
| Flexibility | Limited | High |
| Type Safety | No | Yes (Go) |
| Multi-service | Basic | Excellent |
| Custom Logic | Limited | Full programming |
| Ecosystem | AWS only | AWS + custom |

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is licensed under the MIT License.
