package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Generator struct {
	Runtime      string
	FunctionName string
	OutputDir    string
}

func New(runtime, functionName, outputDir string) *Generator {
	return &Generator{
		Runtime:      runtime,
		FunctionName: functionName,
		OutputDir:    outputDir,
	}
}

func (g *Generator) Generate() error {
	switch g.Runtime {
	case "python":
		return g.generatePython()
	case "nodejs":
		return g.generateNodeJS()
	case "go":
		return g.generateGo()
	default:
		return fmt.Errorf("unsupported runtime: %s", g.Runtime)
	}
}

func (g *Generator) generatePython() error {
	// Create app.py
	appPyContent := fmt.Sprintf(`import json


def lambda_handler(event, context):
    """
    %s Lambda function handler
    
    Args:
        event: Lambda event object
        context: Lambda context object
        
    Returns:
        dict: Response object with statusCode and body
    """
    
    print(f"Received event: {json.dumps(event)}")
    
    # Your function logic here
    response_body = {
        "message": "Hello from %s!",
        "function": "%s",
        "event": event
    }
    
    return {
        "statusCode": 200,
        "headers": {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*"
        },
        "body": json.dumps(response_body)
    }
`, g.FunctionName, g.FunctionName, g.FunctionName)

	if err := g.writeFile("app.py", appPyContent); err != nil {
		return err
	}

	// Create requirements.txt
	requirementsContent := `# Add your Python dependencies here
# Example:
# requests==2.31.0
# boto3==1.26.137
`
	if err := g.writeFile("requirements.txt", requirementsContent); err != nil {
		return err
	}

	// Create CDK infrastructure
	return g.generateCDKInfrastructure("python3.9", "app.lambda_handler")
}

func (g *Generator) generateNodeJS() error {
	// Create index.js
	indexJsContent := fmt.Sprintf(`/**
 * %s Lambda function handler
 * 
 * @param {Object} event - Lambda event object
 * @param {Object} context - Lambda context object
 * @returns {Object} Response object with statusCode and body
 */
exports.handler = async (event, context) => {
    console.log('Received event:', JSON.stringify(event, null, 2));
    
    // Your function logic here
    const responseBody = {
        message: 'Hello from %s!',
        function: '%s',
        event: event
    };
    
    return {
        statusCode: 200,
        headers: {
            'Content-Type': 'application/json',
            'Access-Control-Allow-Origin': '*'
        },
        body: JSON.stringify(responseBody)
    };
};

// For local testing
if (require.main === module) {
    const testEvent = {
        httpMethod: 'GET',
        path: '/test',
        queryStringParameters: null,
        body: null
    };
    
    exports.handler(testEvent, {})
        .then(result => console.log(JSON.stringify(result, null, 2)))
        .catch(error => console.error('Error:', error));
}
`, g.FunctionName, g.FunctionName, g.FunctionName)

	if err := g.writeFile("index.js", indexJsContent); err != nil {
		return err
	}

	// Create package.json
	packageJsonContent := fmt.Sprintf(`{
  "name": "%s",
  "version": "1.0.0",
  "description": "%s Lambda function",
  "main": "index.js",
  "scripts": {
    "test": "node index.js",
    "deploy": "./deploy.sh"
  },
  "dependencies": {
  },
  "devDependencies": {
  },
  "author": "",
  "license": "MIT"
}
`, g.FunctionName, g.FunctionName)

	if err := g.writeFile("package.json", packageJsonContent); err != nil {
		return err
	}

	// Create CDK infrastructure
	return g.generateCDKInfrastructure("nodejs18.x", "index.handler")
}

func (g *Generator) generateGo() error {
	// Create main.go
	mainGoContent := fmt.Sprintf(`package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response represents the Lambda function response
type Response struct {
	Message  string      `+"`json:\"message\"`"+`
	Function string      `+"`json:\"function\"`"+`
	Event    interface{} `+"`json:\"event\"`"+`
}

// HandleRequest handles the Lambda request
func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received event: %%+v", request)

	// Your function logic here
	response := Response{
		Message:  "Hello from %s!",
		Function: "%s",
		Event:    request,
	}

	responseBody, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{}, fmt.Errorf("failed to marshal response: %%w", err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
		},
		Body: string(responseBody),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
`, g.FunctionName, g.FunctionName)

	if err := g.writeFile("main.go", mainGoContent); err != nil {
		return err
	}

	// Create go.mod
	goModContent := fmt.Sprintf(`module %s

go 1.21

require (
	github.com/aws/aws-lambda-go v1.46.0
)
`, g.FunctionName)

	if err := g.writeFile("go.mod", goModContent); err != nil {
		return err
	}

	// Create CDK infrastructure
	return g.generateCDKInfrastructure("go1.x", "main")
}

func (g *Generator) writeFile(filename, content string) error {
	filePath := filepath.Join(g.OutputDir, filename)

	// Create directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, []byte(content), 0644)
}

// generateCDKInfrastructure creates AWS CDK infrastructure code
func (g *Generator) generateCDKInfrastructure(runtime, handler string) error {
	// Create CDK app main file
	cdkMainContent := fmt.Sprintf(`package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type %sStackProps struct {
	awscdk.StackProps
}

func New%sStack(scope constructs.Construct, id string, props *%sStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Lambda function
	lambdaFunction := awslambda.NewFunction(stack, jsii.String("%sFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_%s(),
		Handler: jsii.String("%s"),
		Code:    awslambda.Code_FromAsset(jsii.String(".")),
		FunctionName: jsii.String("%s"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
		MemorySize: jsii.Number(128),
		Environment: &map[string]*string{
			"FUNCTION_NAME": jsii.String("%s"),
		},
	})

	// API Gateway
	api := awsapigateway.NewRestApi(stack, jsii.String("%sApi"), &awsapigateway.RestApiProps{
		RestApiName: jsii.String("%s-api"),
		Description: jsii.String("API for %s Lambda function"),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: jsii.Strings("Content-Type", "X-Amz-Date", "Authorization", "X-Api-Key", "X-Amz-Security-Token"),
		},
	})

	// Lambda integration
	lambdaIntegration := awsapigateway.NewLambdaIntegration(lambdaFunction, &awsapigateway.LambdaIntegrationOptions{
		RequestTemplates: &map[string]*string{
			"application/json": jsii.String("{\"statusCode\": \"200\"}"),
		},
	})

	// Add resource and method
	resource := api.Root().AddResource(jsii.String("%s"), nil)
	resource.AddMethod(jsii.String("ANY"), lambdaIntegration, nil)
	
	// Add proxy resource for catch-all
	proxyResource := resource.AddResource(jsii.String("{proxy+}"), nil)
	proxyResource.AddMethod(jsii.String("ANY"), lambdaIntegration, nil)

	// Outputs
	awscdk.NewCfnOutput(stack, jsii.String("ApiUrl"), &awscdk.CfnOutputProps{
		Value:       api.Url(),
		Description: jsii.String("API Gateway URL"),
	})

	awscdk.NewCfnOutput(stack, jsii.String("FunctionArn"), &awscdk.CfnOutputProps{
		Value:       lambdaFunction.FunctionArn(),
		Description: jsii.String("Lambda Function ARN"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	New%sStack(app, "%sStack", &%sStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}
`,
		capitalizeFirst(g.FunctionName),
		capitalizeFirst(g.FunctionName),
		capitalizeFirst(g.FunctionName),
		capitalizeFirst(g.FunctionName),
		getRuntimeForCDK(runtime),
		handler,
		g.FunctionName,
		g.FunctionName,
		capitalizeFirst(g.FunctionName),
		g.FunctionName,
		g.FunctionName,
		g.FunctionName,
		capitalizeFirst(g.FunctionName),
		capitalizeFirst(g.FunctionName),
		capitalizeFirst(g.FunctionName))

	if err := g.writeFile("infrastructure/main.go", cdkMainContent); err != nil {
		return err
	}

	// Create CDK go.mod
	cdkGoModContent := fmt.Sprintf(`module %s-infrastructure

go 1.21

require (
	github.com/aws/aws-cdk-go/awscdk/v2 v2.161.1
	github.com/aws/constructs-go/constructs/v10 v10.3.0
	github.com/aws/jsii-runtime-go v1.103.1
)
`, g.FunctionName)

	if err := g.writeFile("infrastructure/go.mod", cdkGoModContent); err != nil {
		return err
	}

	// Create CDK app configuration
	cdkJsonContent := `{
  "app": "go run main.go",
  "watch": {
    "include": [
      "**"
    ],
    "exclude": [
      "README.md",
      "cdk*.json",
      "**/*.d.ts",
      "**/*.js",
      "tsconfig.json",
      "package*.json",
      "yarn.lock",
      "node_modules",
      "test"
    ]
  },
  "context": {
    "@aws-cdk/aws-lambda:recognizeLayerVersion": true,
    "@aws-cdk/core:checkSecretUsage": true,
    "@aws-cdk/core:target-partitions": [
      "aws",
      "aws-cn"
    ],
    "@aws-cdk-containers/ecs-service-extensions:enableDefaultLogDriver": true,
    "@aws-cdk/aws-ec2:uniqueImdsv2TemplateName": true,
    "@aws-cdk/aws-ecs:arnFormatIncludesClusterName": true,
    "@aws-cdk/aws-iam:minimizePolicies": true,
    "@aws-cdk/core:validateSnapshotRemovalPolicy": true,
    "@aws-cdk/aws-codepipeline:crossAccountKeyAliasStackSafeResourceName": true,
    "@aws-cdk/aws-s3:createDefaultLoggingPolicy": true,
    "@aws-cdk/aws-sns-subscriptions:restrictSqsDescryption": true,
    "@aws-cdk/aws-apigateway:disableCloudWatchRole": true,
    "@aws-cdk/core:enablePartitionLiterals": true,
    "@aws-cdk/aws-events:eventsTargetQueueSameAccount": true,
    "@aws-cdk/aws-iam:standardizedServicePrincipals": true,
    "@aws-cdk/aws-ecs:disableExplicitDeploymentControllerForCircuitBreaker": true,
    "@aws-cdk/aws-iam:importedRoleStackSafeDefaultPolicyName": true,
    "@aws-cdk/aws-s3:serverAccessLogsUseBucketPolicy": true,
    "@aws-cdk/aws-route53-patters:useCertificate": true,
    "@aws-cdk/customresources:installLatestAwsSdkDefault": false,
    "@aws-cdk/aws-rds:databaseProxyUniqueResourceName": true,
    "@aws-cdk/aws-codedeploy:removeAlarmsFromDeploymentGroup": true,
    "@aws-cdk/aws-apigateway:authorizerChangeDeploymentLogicalId": true,
    "@aws-cdk/aws-ec2:launchTemplateDefaultUserData": true,
    "@aws-cdk/aws-secretsmanager:useAttachedSecretResourcePolicyForSecretTargetAttachments": true,
    "@aws-cdk/aws-redshift:columnId": true,
    "@aws-cdk/aws-stepfunctions-tasks:enableLoggingsfnStateMachine": true,
    "@aws-cdk/aws-ec2:restrictDefaultSecurityGroup": true,
    "@aws-cdk/aws-apigateway:requestValidatorUniqueId": true,
    "@aws-cdk/aws-kms:aliasNameRef": true,
    "@aws-cdk/aws-autoscaling:generateLaunchTemplateInsteadOfLaunchConfig": true,
    "@aws-cdk/core:includePrefixInUniqueNameGeneration": true,
    "@aws-cdk/aws-efs:denyAnonymousAccess": true,
    "@aws-cdk/aws-opensearchservice:enableOpensearchMultiAzWithStandby": true,
    "@aws-cdk/aws-lambda-nodejs:useLatestRuntimeVersion": true,
    "@aws-cdk/aws-efs:mountTargetOrderInsensitiveLogicalId": true,
    "@aws-cdk/aws-rds:auroraClusterChangeScopeOfInstanceParameterGroupWithEachParameters": true,
    "@aws-cdk/aws-appsync:useArnForSourceApiAssociationIdentifier": true,
    "@aws-cdk/aws-rds:preventRenderingDeprecatedCredentials": true,
    "@aws-cdk/aws-codepipeline-actions:useNewDefaultBranchForSourceAction": true
  }
}`

	if err := g.writeFile("infrastructure/cdk.json", cdkJsonContent); err != nil {
		return err
	}

	// Create deployment scripts
	deployScript := fmt.Sprintf(`#!/bin/bash

# %s Deployment Script

set -e

echo "🚀 Deploying %s function..."

# Navigate to infrastructure directory
cd infrastructure

# Install CDK dependencies
echo "📦 Installing CDK dependencies..."
go mod tidy

# Bootstrap CDK (run this once per account/region)
echo "🔧 Bootstrapping CDK..."
cdk bootstrap

# Deploy the stack
echo "📋 Deploying CloudFormation stack..."
cdk deploy --require-approval never

echo "✅ Deployment completed!"
echo ""
echo "🔗 Check the outputs above for your API Gateway URL"
`, g.FunctionName, g.FunctionName)

	if err := g.writeFile("deploy.sh", deployScript); err != nil {
		return err
	}

	// Make deploy script executable
	deployPath := filepath.Join(g.OutputDir, "deploy.sh")
	if err := os.Chmod(deployPath, 0755); err != nil {
		return err
	}

	// Create README for CDK deployment
	readmeContent := fmt.Sprintf(`# %s Function

This serverless function is built with AWS CDK for Go.

## Structure

- app.py/index.js/main.go - Your function code
- infrastructure/ - CDK infrastructure code
- deploy.sh - Deployment script

## Prerequisites

- AWS CLI configured
- AWS CDK CLI installed (npm install -g aws-cdk)
- Go 1.21+

## Deployment

1. Make sure your AWS credentials are configured:
   aws configure

2. Run the deployment script:
   ./deploy.sh

   Or manually:
   cd infrastructure
   go mod tidy
   cdk bootstrap
   cdk deploy

## Development

- Modify your function code in the main file
- Update infrastructure in infrastructure/main.go
- Redeploy with ./deploy.sh

## Clean Up

To delete the stack:
cd infrastructure
cdk destroy
`, g.FunctionName)

	return g.writeFile("README.md", readmeContent)
}

// Helper functions
func capitalizeFirst(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func getRuntimeForCDK(runtime string) string {
	switch runtime {
	case "python3.9":
		return "PYTHON_3_9"
	case "nodejs18.x":
		return "NODEJS_18_X"
	case "go1.x":
		return "GO_1_X"
	default:
		return "PYTHON_3_9"
	}
}
