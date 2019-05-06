import json


def lambda_handler(event, context):
    """Sample pure Lambda function

    Arguments:
        event LambdaEvent -- Lambda Event received from Invoke API
        context LambdaContext -- Lambda Context runtime methods and attributes

    Returns:
        dict -- {'statusCode': int, 'body': dict}
    """

    return {
        "statusCode": 200,
        "body": json.dumps({
            'message': 'API deployed',
        })
    }
