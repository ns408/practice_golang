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
  if assignment-statement; condition {
  }

  //Idiomatic Go
  - avoid unnecessary branches and indentation of code.
  - return as early as possible.
  `)

	num := 10
	if num%2 == 0 { // Is the number divisble by 2 with 0 remainder.
		fmt.Println("The number", num, "is even")
		return
	}
	fmt.Println("The number", num, "is odd")

}
