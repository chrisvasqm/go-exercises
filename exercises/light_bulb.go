package exercises

import "fmt"

type LightBulb struct {
	isOn bool
}

func (lightbulb *LightBulb) TurnOn() {
	lightbulb.isOn = true
	fmt.Println("Lightbulb is ON")
}

func (lightbulb *LightBulb) TurnOff() {
	lightbulb.isOn = false
	fmt.Println("Lightbulb is OFF")
}

func (lightbulb *LightBulb) IsOn() bool {
	return lightbulb.isOn
}