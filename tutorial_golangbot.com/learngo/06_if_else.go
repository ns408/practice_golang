package main

import "fmt"

func main() {

	fmt.Println(`
  //Notes:
  //'if syntax':

  if condition {
  }

  if condition {
  } else {
  }

  if condition1 {
  } else if condition2 {
  } else {
  }

  //If with assignment
  //variant which includes an optional shorthand assignment statement.
  //it is executed before the condition is evaluated
  if assignment-statement; condition {
  }

  //Idiomatic Go
  - avoid unnecessary branches and indentation of code.
  - return as early as possible.

  //else statement should start in the same line after the closing curly brace
  //} of the if statement otherwise, error:
  //syntax error: unexpected else, expecting }

  //Context: Go inserts semicolons automatically and the way it does it
  //requires else to end on the same linke as the closing brace.
  //Reference: https://golang.org/ref/spec#Semicolons.
  `)

	num := 9
	if num%2 == 0 { // Is the number divisble by 2 with 0 remainder.
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd\n")

	num = 99 // using = instead of := (shorthand) because otherwise following issue
	//no new variables on left side of :=
	if num <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}

	// if with assignment
	// here 'num' variable will be accessible only within the scope of if-else
	// blocks.
	if num = 50; num <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}

}
