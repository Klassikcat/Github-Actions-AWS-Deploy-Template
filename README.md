# GitHub Actions AWS ECS Deployment Template

## Overview
This repository provides a comprehensive GitHub Actions workflow template for deploying applications to AWS ECS (Elastic Container Service). It includes support for Blue/Green deployments through AWS CodeDeploy and automated task definition updates.

## Features
- 🚀 Automated AWS ECS service deployment
- 🔄 Blue/Green deployment support via CodeDeploy
- 🔧 Automatic Task Definition updates
- 📊 Deployment status monitoring
- 🔐 Secure AWS authentication using STS assume role
- 🐳 Multi-container support with sidecar patterns
- 🏗️ Flexible build and deployment configurations

## Prerequisites
- AWS Account with appropriate permissions
- GitHub repository with GitHub Actions enabled
- Docker for local development and testing

## Quick Start
1. Fork this repository or use it as a template
2. Create Identity Provider. in AWS: **[[Guide](https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services)]**
3. Create IAM Role in your AWS Account:
   ```json
   {
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Principal": {
                "Federated": "arn:aws:iam::<ACCOUNT-ID-HERE>:oidc-provider/token.actions.githubusercontent.com"
            },
            "Action": "sts:AssumeRoleWithWebIdentity",
            "Condition": {
                "StringEquals": {
                    "token.actions.githubusercontent.com:aud": "sts.amazonaws.com"
                },
                "StringLike": {
                    "token.actions.githubusercontent.com:sub": "repo:<ORG-OR-USER-NAME-HERE>/*"
                    }
                }
            }
        ]
    }
    ```
   

3. Copy the `.github/workflows/example-workflow.yaml` to your project
4. Customize the workflow configuration for your needs
5. Push your changes to trigger the workflow

## Workflow Structure
- `build.yaml`: Handles Docker image building and pushing to ECR
- `deploy.yaml`: Manages ECS service deployment and updates
- `deploy_lambda.yaml`: Optional Lambda function deployment
- `cancel.yaml`: Workflow cancellation handling
- `run-credential-searcher.yaml`: Security scanning for credentials

## Docker Support
The repository includes two Dockerfile templates:
- `server.Dockerfile`: Main application container
- `sidecar.Dockerfile`: Sidecar container for additional services

## Security
- ⚠️ **Important**: Do not use `AWS_ACCESS_KEY` and `AWS_SECRET_ACCESS_KEY` for GitHub Actions AWS Authentication. Use `sts-assume-role` instead.
- Regular security scanning for exposed credentials
- Secure secret management through GitHub Secrets

## Contributing
We welcome contributions! Please see our [Contributing Guidelines](.github/CONTRIBUTING.md) for details.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support
If you encounter any issues or have questions:
1. Check the [existing issues](https://github.com/yourusername/github-actions-aws-deploy-template-docker/issues)
2. Create a new issue using our [issue templates](.github/ISSUE_TEMPLATE/)
3. Review our [documentation](docs/)

## Acknowledgments
- AWS ECS Team
- GitHub Actions Team
- All contributors and users of this template
