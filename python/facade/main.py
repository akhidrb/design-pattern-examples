class CarEngineFacade(object):
    def on(self):
        self.fuelInjector = FuelInjector()
        self.fuelInjector.on()
        self.fuelInjector.inject()

        self.starter = Starter()
        self.starter.start()


class FuelInjector(object):
    def on(self):
        print("FuelInjector turned on")

    def inject(self):
        print("FuelInjector injects fuel")


class Starter(object):
    def start(self):
        print("Starter starts")


if __name__ == "__main__":
    carEngine = CarEngineFacade()
    carEngine.on()