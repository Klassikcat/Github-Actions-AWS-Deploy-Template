def lambda_handler(event, context):
    print(context)
    return {
        'statusCode': 200,
        'body': 'Hello, World!'
    }