package main

import "fmt"
import "math"

// Notes:
// Automatic initialisation of a variable to **zero value** if not assigned a value.

func main() {

	// Declaring a variable and assign a value
	var age int // declare variable named `age` of type `int`.
	fmt.Println("What was the age - ", age)

	age = 29 // assigned a value
	fmt.Println("What was the age - ", age)

	age = 2 // assigned a value
	fmt.Println("What was the age - ", age)

	// Declaring variable with an initial value
	var num int = 100 // declare variable with initial value
	fmt.Println("The number was - ", num)

	// Type inference
	// Go automatically infers the type of variable if it has an initial value
	var name = "this"
	fmt.Println("My name is ", name)
	var name1 = 'x'
	fmt.Println("My name is ", name1)

	// Multiple variable declaration
	var width, height int = 100, 50
	fmt.Println("width is", width, "height is", height)

	// declare variables belonging to different types in a single statement
	var (
		name2   = "saywhatnow"
		age1    = 90
		height1 int
	)
	fmt.Println("Name is", name2, "\nAge is", age1, "\nHeight is", height1)

	// Short hand declaration
	// name := initialvalue

	// declare `count` - go will infer the type
	count := 10
	fmt.Println("Count =", count)

	// Short hand declaration for multiple variables
	var_name, var_age := "batman", 50
	fmt.Println("My name is", var_name, "; Age", var_age)

	// Short hand syntax requires at least one of the vars to be newly declared
	vara, varb := 20, 30
	varb, varc := 10, 40
	fmt.Println("vara", vara, ";varb", varb, ";varc", varc)

	// Assign a run time computed value to a variable
	var_d, var_e := 50.0, 51.0
	var_c := math.Min(var_d, var_e)
	fmt.Println("Minimum value is", var_c)
	// Show me variable type genie
	fmt.Printf("%T\n", var_c)
	fmt.Printf("%V\n", var_c)
	fmt.Printf("%0.2f\n", var_c)
	// math.Min function requires float as arguments
	// cannot use var_e (type int) as type float64 in argument to math.Min

	// Note:
	//
	// Can't redeclare a variable in the same block
	// var num int = 100 // declare variable with initial value
	// var num int = 100 // declare variable with initial value
	// ./variables.go:21:7: age redeclared in this block
	//       previous declaration at ./variables.go:11:6

	// Can't assign a value of a different type to a variable already declared
	// var num int = 100 // declare variable with initial value
	// var num = '100'
	// ./variables.go:29:7: num redeclared in this block
	//          previous declaration at ./variables.go:21:7
	//  ./variables.go:29:13: invalid character literal (more than one character)

	// Use double quotes ""
	// var name = 'A_NAME' will result in the following error
	// ./variables.go:27:14: invalid character literal (more than one character)

	// While using short hand for declaring two variables, one has to assign values
	// ./01_variables.go:51:21: assignment mismatch: 2 variables but 1 values

	// At least one new variable is needed for short hand multi var declaration
	// ./01_variables.go:57:14: no new variables on left side of :=
}
