package bdd.questions;

import io.restassured.response.Response;
import net.serenitybdd.screenplay.Question;

public class HttpResponse {

    public static Question<Integer> statusCodeFor(final Response response) {
        return actor -> response.statusCode();
    }

    public static Question<String> message(final Response response) {
        return actor -> response.jsonPath().get("message");
    }
}
