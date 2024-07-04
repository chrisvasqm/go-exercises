package main

import (
	"fmt"
	g "goexercises/greet"
	"goexercises/models"
)

func main() {
	chris := models.Person{Name: "Chris", Age: 29}
	fmt.Println(g.Greet(chris.Name))
}	
