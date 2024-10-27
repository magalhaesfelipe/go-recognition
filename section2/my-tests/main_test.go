package main

import (
	"fmt"
	"testing"
)

// USING A TABLE TEST

var someTests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-250", 500.0, 2.0, 250.0, false},
	{"expect-3", 9.0, 3.0, 3.0, false},
	{"expect-2.875", -23.0, -8.0, 2.875, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range someTests {
		divisionResult, err := divideNumbers(tt.dividend, tt.divisor)
		fmt.Println(divisionResult)

		// Check if an error was expected		
		if tt.isErr {
			if err == nil { // with invalid data, 'err' cannot be 'nil', so we fail the test
				t.Errorf("Test -> %s: Expected an error but did not get one", tt.name)
			} 
		} else {
				// No error expected
				if err != nil { // with valid data, 'err' must be 'nil', so we fail the test
					t.Errorf("Test -> %s: Did not expect an error but got one: %s", tt.name, err.Error())
					} else if divisionResult != tt.expected { // Only check result if no error
					t.Errorf("Test -> %s: Expected %f and got %f", tt.name, tt.expected, divisionResult)
			  } 
		}
	}
}

// // Writing one function for every test

// // A function to test receives an argument/parameter 't' which is a pointer/reference to 'T' from the 'testing' Package

// // TESTING FOR A VALID DIVISION
// func TestValidDivision(t *testing.T) {
// 	_, err := divideNumbers(2.0, 1.0) // 2 / 1 should not generate an error because it's a valid operation, if it throws an error though, it will fail this test
// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// // TESTING FOR AN INVALID DIVISION
// func TestInvalidDivision(t *testing.T) {
// 	_, err := divideNumbers(2, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }
