package bdd.tasks;

import net.serenitybdd.rest.SerenityRest;
import net.serenitybdd.screenplay.Actor;
import net.serenitybdd.screenplay.Task;
import net.serenitybdd.screenplay.rest.interactions.Get;
import net.thucydides.core.annotations.Step;

import static net.serenitybdd.screenplay.Tasks.instrumented;

public class HealthEndpoint implements Task {

    public static HealthEndpoint get() {
        return instrumented(HealthEndpoint.class);
    }

    @Override
    @Step("{0} makes a HTTP GET call")
    public <T extends Actor> void performAs(final T actor) {
        actor.attemptsTo(
                Get.resource("/__healthcheck")
        );
        actor.remember("response", SerenityRest.lastResponse());
    }
}
