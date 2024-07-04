package main

import (
	"fmt"
	"goexercises/exercises"
	"goexercises/greet"
	"goexercises/models"
)

func main() {
	chris := models.Person{Name: "Chris", Age: 29}
	fmt.Println(greet.Greet(chris.Name))

	for i := 1; i <= 20; i++ {
		fmt.Println(exercises.FizzBuzz(i))
	}

	lightbulb := exercises.Lightbulb{}
	fmt.Println(lightbulb.IsOn())
	lightbulb.TurnOn()
	lightbulb.TurnOff()
	fmt.Println(lightbulb.IsOn())
}	
