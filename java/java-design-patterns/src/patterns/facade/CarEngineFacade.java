package patterns.facade;

import patterns.facade.subcomponents.AirFlowController;
import patterns.facade.subcomponents.FuelInjector;
import patterns.facade.subcomponents.Starter;

public class CarEngineFacade {

    private AirFlowController airFlowController;
    private FuelInjector fuelInjector;
    private Starter starter;

    public CarEngineFacade() {
        airFlowController = new AirFlowController();
        fuelInjector = new FuelInjector();
        starter = new Starter();
    }

    public void on() {
        airFlowController.takeAir();
        fuelInjector.on();
        fuelInjector.inject();
        starter.start();
    }


}
