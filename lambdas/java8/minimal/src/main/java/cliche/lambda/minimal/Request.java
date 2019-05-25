package cliche.lambda.minimal;

import java.util.Map;
import java.util.Optional;

public class Request {
    private Map<String, String> queryStringParameters;

    public Request() {

    }

    Map<String, String> getQueryStringParameters() {
        return this.queryStringParameters;
    }

    Optional<String> getHomeTeam() {
        return Optional.ofNullable(this.queryStringParameters.get("HomeTeam"));
    }

    Optional<String> getAwayTeam() {
        return Optional.ofNullable(this.queryStringParameters.get("AwayTeam"));
    }

    public void setQueryStringParameters(final Map<String, String> queryStringParameters) {
        this.queryStringParameters = queryStringParameters;
    }
}
