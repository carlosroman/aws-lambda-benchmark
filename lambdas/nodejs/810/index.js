'use strict';

console.log('Loading function');
const dynamodb = require('aws-sdk/clients/dynamodb');
const dynamodbConfig = {
    region: process.env.TABLE_REGION,
    endpoint: process.env.ENDPOINT_OVERRIDE,
};
const docClient = new dynamodb.DocumentClient(dynamodbConfig);

exports.handler = async (event) => {
	
	console.log(`Processing Lambda request ${event.requestContext.requestId}`);
	let data;
	// if (!('HomeTeam' in event.queryStringParameters)) {
	if (event.queryStringParameters.HomeTeam === undefined) {
        return {
		    statusCode: 400,
		    body: 'Missing param "HomeTeam"',
		    headers: {},
		};
	}
	// if (!('AwayTeam' in event.queryStringParameters)) {
	if (event.queryStringParameters.AwayTeam === undefined) {
        return {
		    statusCode: 400,
		    body: 'Missing param "AwayTeam"',
		    headers: {},
		};
	}

	const params = {
	    TableName : process.env.TABLE_NAME,
	    Key:{
	        "HomeTeam": event.queryStringParameters.HomeTeam,
	        "AwayTeam": event.queryStringParameters.AwayTeam,
	    }
	};

	try {
		console.log('DynamoDB Get params:', JSON.stringify(params, null, 2));
        data = await docClient.get(params).promise();
    }
    catch (err) {
        console.log(err);
        return {
		    statusCode: 500,
		    body: JSON.stringify(err),
		    headers: {},
		};
	}
	
	return {
		statusCode: 200,
		body: JSON.stringify(data.Item),
		headers: {
			// "Content-Type": "application/json",
		},
	};
}
