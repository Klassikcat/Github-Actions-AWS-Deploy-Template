# Instruction(Copilot)

## PoorServerless

### Goal

Deploy app.py and corresponding serverless functions without handling template.yaml(SAM Template)

### Elements

#### Function

Your function should create following:

- app.py(with lambda_handler(event, context) function) or corresponding node.js, other serverless handler and file
- spec.yaml that contains all necessary informations

##### spec.yaml

```yaml
name: function-name
transforms:
    - AWS::Serverless-2016-10-31
    - ${additional-transforms}
spec:
    runtime: ${funciton-runtime} // default = python3.11
    memory: ${memory-size} // default = 512
    timeout: ${timeout-seconds} // default = 30
layers:
    - ${layer-name}
roles:
    - ${role-name}
policies:
    - ${policy-names}
invoke-permissions:
    - type: ${permission-type}
      principal: ${principal-name}
      source-arn: ${source-arn}
envs:
  ${env-name}: ${env-value}
  ${env-name}: ${env-value}
```

**Example 1: Uses Externel Role**

Let's suppose there function that can be expressed by SAM Template like this
<details>
<summary>Example SAM Template</summary>

```yaml 
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

Resources:
  Function:
    Type: AWS::Serverless::Function
    Properties:
      FunctionName: function-name
      CodeUri: ./
      Runtime: python3.11
      Handler: app.lambda_handler
      Layers:
        - Layer1
        - Layer2
      MemorySize: 1024
      Timeout: 30
      Environment:
        Variables:
          env: !Ref Stage
          TZ: Asia/Seoul
      Roles:
        - arn:aws:iam::123456789012:role/LambdaExecutionRole1
        - arn:aws:iam::123456789012:role/LambdaExecutionRole2
        - arn:aws:iam::123456789012:role/LambdaExecutionRole3
  LambdaInvokePermission:
    Type: AWS::Lambda::Permission
    Properties:
      Action: lambda:InvokeFunction
      FunctionName: !GetAtt Function.Arn
      Principal: apigateway.amazonaws.com
      SourceArn: !Sub "arn:aws:execute-api:${AWS::Region}:${AWS::AccountId}:api-gateway-resource-arn/*/METHOD/PATH"
  LambdaInvokePermissionSub:
    Type: AWS::Lambda::Permission
    Properties:
      Action: lambda:InvokeFunction
      FunctionName: !GetAtt Function.Arn
      Principal: apigateway.amazonaws.com
      SourceArn: !Sub "arn:aws:execute-api:${AWS::Region}:${AWS::AccountId}:api-gateway-resource-arn/*/METHOD/PATH/*"
```
</details>

This should be look like spec.yaml like below

```yaml
name: function-name
spec:
    - runtime: python3.11
    - memory: 1024
    - timeout: 30
layers:
    - Layer1
    - Layer2
roles:
    - arn:aws:iam::123456789012:role/LambdaExecutionRole1
    - arn:aws:iam::123456789012:role/LambdaExecutionRole2
permssions:
    - name: LambdaInvokePermission
      principal: apigateway.amazonaws.com
      source-arn: api-gateway-resource-arn/*/METHOD/PATH
    - name: LambdaInvokePermssionNested
      principal: apigateway.amazonaws.com
      source-arn: api-gateway-resource-arn/*/METHOD/PATH/*
envs:
  stage: !Ref Stage
  TZ: Asia/Seoul
```

**Example 2: Attach policies directly to function**

Let's suppose there is function that can be expressed by SAM Template like this

<details>
<summary>Example SAM Template</summary>

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: 
  - AWS::Serverless-2016-10-31
  - Name: DatadogServerless
    Parameters:
      stackName: !Ref "AWS::StackName"
      pythonLayerVersion: 111
      extensionLayerVersion: 81
      site: "datadoghq.com"
      apiKeySecretArn: "arn:aws:secretsmanager:ap-northeast-2:1111111:secret:DdApiKeySecret-SAdfSAe"

Resources:
  Function:
    Type: AWS::Serverless::Function
    Properties:
      FunctionName: function-name
      CodeUri: ./
      Runtime: python3.11
      Handler: app.lambda_handler
      Layers:
        - Layer1
        - Layer2
      MemorySize: 1024
      Timeout: 30
      Environment:
        Variables:
          env: !Ref Stage
          TZ: Asia/Seoul
      Policies:
        - AWSLambdaBasicExecutionRole
        - AWSLambdaVPCAccessExecutionRole
        - SecretsManagerReadWrite
        - CloudWatchFullAccess
  LambdaExecutionPolicy:
    Type: AWS::IAM::Policy
    Properties:
      PolicyName: LambdaExecutionPolicy
      PolicyDocument:
        Version: '2012-10-17'
        Statement:
          - Effect: Allow
            Action:
              - logs:CreateLogGroup
              - logs:CreateLogStream
              - logs:PutLogEvents
            Resource: arn:aws:logs:*:*:*
  LambdaInvokePermission:
    Type: AWS::Lambda::Permission
    Properties:
      Action: lambda:InvokeFunction
      FunctionName: !GetAtt Function.Arn
      Principal: apigateway.amazonaws.com
      SourceArn: !Sub "arn:aws:execute-api:${AWS::Region}:${AWS::AccountId}:api-gateway-resource-arn/*/METHOD/PATH"
  LambdaInvokePermissionSub:
    Type: AWS::Lambda::Permission
    Properties:
      Action: lambda:InvokeFunction
      FunctionName: !GetAtt Function.Arn
      Principal: apigateway.amazonaws.com
      SourceArn: !Sub "arn:aws:execute-api:${AWS::Region}:${AWS::AccountId}:api-gateway-resource-arn/*/METHOD/PATH/*"
```
</details>

Then, this should be look like spec.yaml like below

```yaml
name: function-name
transforms:
  - Name: DatadogServerless
    Parameters:
      stackName: !Ref "AWS::StackName"
      pythonLayerVersion: 111
      extensionLayerVersion: 81
      site: "datadoghq.com"
      apiKeySecretArn: "arn:aws:secretsmanager:ap-northeast-2:1111111:secret:DdApiKeySecret-SAdfSAe"
spec:
  - runtime: python3.11
  - memory: 1024
  - timeout: 30
layers:
  - Layer1
  - Layer2
policies:
  - AWSLambdaBasicExecutionRole
  - AWSLambdaVPCAccessExecutionRole
  - SecretsManagerReadWrite
  - CloudWatchFullAccess
permssions:
    - name: LambdaInvokePermission
      principal: apigateway.amazonaws.com
      source-arn: api-gateway-resource-arn/*/METHOD/PATH
    - name: LambdaInvokePermssionNested
      principal: apigateway.amazonaws.com
      source-arn: api-gateway-resource-arn/*/METHOD/PATH/*
```
#### Layers

<!-- TODO: Add layer generator -->