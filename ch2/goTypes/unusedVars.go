package goTypes

import "fmt"

// Every declared local variable must be read
// It's a compile-time error to declare a local variable and to not red its value

// The Go compiler won't stop you from creating unread package-level variables. This is one more reason you should avoid creating package-level variables
var Yes string = "hello" // UNUSED

func Test() {

	// As long as a variable is read once, the compiler won't complain, even if there are writes to the variable that are never read

	// You could call it 'unused assignments' (the compiler and 'go vet' do not catch the unused assignments)
	a := 20 // this assignment ins't read
	a = 389
	fmt.Println(a)
	a = 40 // this assignment ins't read

}

func Test2() {

	// The Go compiler allow you to create unread constants with const
	const whatever int = 20 // UNUSED BUT NO ERROR

	// Constants is Go are calculated at compile time. This makes them easy to eliminate: if a constant isn't used, it is simply not included in the compiled binary

}
