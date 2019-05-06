package bdd.stepdefinitions;


import bdd.tasks.HealthEndpoint;
import cucumber.api.java.en.Given;
import cucumber.api.java.en.Then;
import cucumber.api.java.en.When;
import net.serenitybdd.screenplay.Actor;
import net.serenitybdd.screenplay.rest.abilities.CallAnApi;

import static net.serenitybdd.screenplay.rest.questions.ResponseConsequence.seeThatResponse;
import static org.hamcrest.Matchers.equalTo;
public class StatusStepdefs {

    private Actor theHealthChecker;

    @Given("the (.*) wants to check the API")
    public void the_actor_wants_to_check_the_API(String actor) {
        this.theHealthChecker = Actor.named(actor).whoCan(CallAnApi.at("http://localhost:3000"));
    }

    @When("they check the application status")
    public void they_check_the_application_status() {
        this.theHealthChecker.attemptsTo(HealthEndpoint.get());
    }

    @Then("the status code should be {int}")
    public void the_status_code_should_be(final int statusCode) {
        this.theHealthChecker.should(
                seeThatResponse("The correct status code was returned",
                        response -> response.statusCode(statusCode))
        );
//        final Response theResponse = this.theHealthChecker.recall("response");
//        this.theHealthChecker.should(seeThat(HttpResponse.statusCodeFor(theResponse), equalTo(statusCode)));
    }

    @Then("the API should return the message {string}")
    public void then_the_API_should_return_the_message(final String message) {
//        final Response theResponse = this.theHealthChecker.recall("response");
//        this.theHealthChecker.should(
//                seeThat(HttpResponse.message(theResponse), equalTo(message)));

        this.theHealthChecker.should(
                seeThatResponse("The correct message was returned",
                        response -> response.body("message", equalTo(message)))
        );
    }
}
