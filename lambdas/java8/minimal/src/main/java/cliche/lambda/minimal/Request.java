package cliche.lambda.minimal;

import java.util.Map;
import java.util.Optional;

public class Request {
    static final String HOME_TEAM = "HomeTeam";
    static final String AWAY_TEAM = "AwayTeam";
    private Map<String, String> queryStringParameters;

    public Request() {

    }

    Map<String, String> getQueryStringParameters() {
        return this.queryStringParameters;
    }

    Optional<String> getHomeTeam() {
        return Optional.ofNullable(this.queryStringParameters.get(HOME_TEAM));
    }

    Optional<String> getAwayTeam() {
        return Optional.ofNullable(this.queryStringParameters.get(AWAY_TEAM));
    }

    public void setQueryStringParameters(final Map<String, String> queryStringParameters) {
        this.queryStringParameters = queryStringParameters;
    }
}
