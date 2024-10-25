package main

import "fmt"

// Interfaces define functions that need to be implemented by
// the variables instantiated using the interface/type satisfies
type Human interface {
	DrinkWater() string
	BrushTeeth() string
}

type Mangaka struct {
	Name      string
	WorkHours int
}

type Programmer struct {
	Name      string
	WorkHours int
}

type Astronaut struct {
	Name      string
	WorkHours int
	Home      string
}

func main() {
	


 // Using shortcut to create the variable
	mangaka := Mangaka{
		Name:      "Akira Toriyama",
		WorkHours: 11,
	}

	programmer := Programmer{
		Name:      "Felipe",
		WorkHours: 10,
	}

	// Manually creating the variable
	var astronaut Astronaut

	// Initialize the fields of the struct
	astronaut.Name = "James"
	astronaut.WorkHours = 24
	astronaut.Home = "Space"


	// The function now accepts 3 different types because they
	// satisfy the interface requirements  
	ShowHumanHabits(&mangaka)
	ShowHumanHabits(&programmer)
	ShowHumanHabits(&astronaut)
}

func ShowHumanHabits(h Human) {
	fmt.Println("This human driks water: ", h.DrinkWater(), ", and Brush his teeth:", h.BrushTeeth(), h)
}

// Implementing the 'Human' interface for Mangaka type-struct
func (x *Mangaka) DrinkWater() string { // Passing a pointer/reference is a good practice
	return "4 times a day"
}

func (x *Mangaka) BrushTeeth() string {
	return "2 times a day"
}

// Implementing the 'Human' interface for Programmer type-struct
func (*Programmer) DrinkWater() string {
	return "7 times a day"
}

func (*Programmer) BrushTeeth() string {
	return "3 times a day"
}

// Implementing the 'Human' interface for Astronaut type-struct
func (*Astronaut) DrinkWater() string {
	return "9 time a day"
}

func (*Astronaut) BrushTeeth() string {
	return "3 times a day"
}
