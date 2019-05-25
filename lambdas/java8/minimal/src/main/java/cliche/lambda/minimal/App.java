package cliche.lambda.minimal;

import com.amazonaws.client.builder.AwsClientBuilder;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDB;
import com.amazonaws.services.dynamodbv2.AmazonDynamoDBClientBuilder;
import com.amazonaws.services.dynamodbv2.document.DynamoDB;
import com.amazonaws.services.dynamodbv2.document.Item;
import com.amazonaws.services.dynamodbv2.document.Table;
import com.amazonaws.services.lambda.runtime.Context;
import com.amazonaws.services.lambda.runtime.LambdaLogger;
import com.amazonaws.services.lambda.runtime.RequestHandler;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

//import org.apache.logging.log4j.LogManager;
//import org.apache.logging.log4j.Logger;

public class App implements RequestHandler<Request, Object> {
//    private static final Logger logger = LogManager.getLogger(App.class);

    private static final AmazonDynamoDB dynamoClient;
    private static final DynamoDB docClient;

    static {
        dynamoClient = AmazonDynamoDBClientBuilder.standard()
                .withEndpointConfiguration(
                        new AwsClientBuilder.EndpointConfiguration(System.getenv("ENDPOINT_OVERRIDE"), System.getenv("TABLE_REGION")))
                .build();
        docClient = new DynamoDB(dynamoClient);
    }

    @Override
    public Object handleRequest(final Request input, final Context context) {
//        logger.info("Processing Lambda request {}", context.getAwsRequestId());
        final LambdaLogger logger = context.getLogger();
        logger.log(String.format("Processing Lambda request %s\n", context.getAwsRequestId()));
        final Map<String, String> headers = new HashMap<>();
        headers.put("Content-Type", "application/json");

        final Optional<String> homeTeam = input.getHomeTeam();
        if (!homeTeam.isPresent()) {
            return new GatewayResponse("Missing param \"HomeTeam\"", headers, 400);
        }
        final Optional<String> awayTeam = input.getAwayTeam();
        if (!awayTeam.isPresent()) {
            return new GatewayResponse("Missing param \"AwayTeam\"", headers, 400);
        }


        final Table table = docClient.getTable(System.getenv("TABLE_NAME"));
        final Item item = table.getItem("HomeTeam", homeTeam.get(), "AwayTeam", awayTeam.get());
        if (item == null) {
            return new GatewayResponse("fixture not found", headers, 404);
        }
        return new GatewayResponse(item.toJSONPretty(), headers, 200);
    }

}

// https://docs.aws.amazon.com/lambda/latest/dg/java-tracing.html
