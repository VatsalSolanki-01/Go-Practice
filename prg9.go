package main

import "fmt"

type Appliance interface {
	TurnOn()
}

func ApplianceTurnOn(ap Appliance) {
	ap.TurnOn()
}

type Fan struct {
}

func (f Fan) TurnOn() {
	fmt.Println("fan is now turnedon")
}

type Light struct {
}

func (l Light) TurnOn() {
	fmt.Println("Light is now tunedon")
}

type Ac struct {
}

func (a Ac) TurnOn() {
	fmt.Println("Ac is not turnedon")
}
