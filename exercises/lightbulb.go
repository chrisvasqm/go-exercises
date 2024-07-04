package exercises

import "fmt"

type Lightbulb struct {
	isOn bool
}

func (lightbulb *Lightbulb) TurnOn() {
	lightbulb.isOn = true
	fmt.Println("Lightbulb is ON")
}

func (lightbulb *Lightbulb) TurnOff() {
	lightbulb.isOn = false
	fmt.Println("Lightbulb is OFF")
}

func (lightbulb *Lightbulb) IsOn() bool {
	return lightbulb.isOn
}