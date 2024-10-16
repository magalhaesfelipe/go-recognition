package structs

import "fmt"

type customType struct {
	FirstName string
}

// Assigning function to the 'customType' struct 
func (*customType) changeMyDestiny() string {
	return "you changed your destiny"
}

// (*customType) --> This is called a receiver, it ties this function to my struct/type, because it pointing * to the struct customType  (*customType)

// This has access to the instantiated variable of our customType,
//  'x'(or whatever other name), so we can access its properties
func (x *customType) printFirstName() string {
	return x.FirstName
}



func Whatever() {
	var one customType
	one.FirstName = "Aizen"

	two := customType{
		FirstName: "Ichigo",
	}

	// Calling the function associated with our customType
	fmt.Println("What happened?", one.changeMyDestiny())

	// Using another associated function of out customType to access is properties and print it
	fmt.Println("This is printed with the function of the 'CustomType':", two.printFirstName())
	
	// Noramally accessing the variable properties of our customType
	fmt.Println("My first variable is:", one.FirstName)
	fmt.Println("My second variable is:", two.FirstName)
}
