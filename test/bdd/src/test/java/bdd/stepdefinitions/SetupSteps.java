package bdd.stepdefinitions;
import cucumber.api.java.Before;
import net.serenitybdd.screenplay.actors.OnStage;
import net.serenitybdd.screenplay.actors.OnlineCast;
public class SetupSteps {


    @Before()
    public void setStage() {
        OnStage.setTheStage(new OnlineCast());
    }
}
