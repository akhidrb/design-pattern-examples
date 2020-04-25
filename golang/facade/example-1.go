package main

import "fmt"

type FuelInjector struct{}

func (f *FuelInjector) on() {
	fmt.Println("FuelInjector is on")
}

func (f *FuelInjector) inject() {
	fmt.Println("FuelInjector injected fuel")
}

type AirController struct {}

func (a *AirController) takeAir() {
	fmt.Println("AirController takes Air")
}

type Starter struct {}

func (s *Starter) on() {
	fmt.Println("Starter is on")
}

type CarEngine struct {
	fuelInjector *FuelInjector
	airController *AirController
	starter *Starter
}

func (c *CarEngine) On() {
	c.fuelInjector.on()
	c.fuelInjector.inject()
	c.airController.takeAir()
	c.starter.on()
}


func main() {
	carEngine := CarEngine{}
	carEngine.On()
}
