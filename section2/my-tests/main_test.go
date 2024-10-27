package main

import "testing"

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
}

func TestDivision(t *testing.T) {
	for _, tt := range someTests {
		got, err := divideNumbers(tt.dividend, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Errorf("Test %s: Expected an error but did not get one", tt.name)
			} else {
				if err != nil {
					t.Errorf("Test %s: Did not expect an error but got one: %s", tt.name, err.Error())
				}
				if got != tt.expected {
					t.Errorf("Test %s: Expected %f and got %f", tt.name,tt.expected, got)
				}
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

// // TESTING FOR A INVALID DIVISION
// func TestInvalidDivision(t *testing.T) {
// 	_, err := divideNumbers(2, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }
