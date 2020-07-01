package main

import "fmt"

func multiply(param_a, param_b int) int {
	var result = param_a * param_b
	return result
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps_named(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return // no need to explicitly return value since 'area' and 'perimeter' are automatically returned.
}

func main() {
	fmt.Println(`
##############
  //function declaration
  func functionname(parametername type) returntype {
    function body
  }

  //parameters and return type are optional.
  func functionname() {
  }

  //multiply two values
  func multiply(param_a int, param_b int) int {
    var result = param_a * param_b
    return result
  }

  //for consecutive parameters of the same type, it's enough to write them once at the end
  func multiply(param_a, param_b int) int {
    var result = param_a * param_b
    return result
  }

  //calling the function
  multiply(10, 10)

  //Multiple return values from a function
  func rectProps(length, width float64) (float64, float64) {
    var area = length * width
    var perimeter = (length + width) * 2
    return area, perimeter
  }

  //Named return values
  // If return value is named, it can be considered as being delcared as a variable in the first line of the function.
  func rectProps_named(length, width float64) (area, perimeter float64) {
    area = length * width
    perimeter = (length + width) * 2
    return // no explicit return value since 'area' and 'perimeter' are automatically returned.

  //Blank Identifier
  // _ is the blank identifier in Go.
  // can be used in place of any value of any type.
  // can be used to discard a value
  area, _ := rectProps(10.8, 5.6) // perimeter will be discarded here
  }
##############
  `)

	fmt.Println("Call multiply function :", multiply(10, 10))

	//Another way
	num_a, num_b := 11, 11
	result := multiply(num_a, num_b)
	fmt.Printf("Call multiply function 2nd time: %v\nLook at the code to see how type inference and short hand for variable assignment has been used.\n\n", result)

	//call rectProps
	area, perimeter := rectProps(10, 5)
	fmt.Printf("Rectangle\nArea: %f, Perimeter %f\n\n", area, perimeter)

	area, perimeter = rectProps_named(13, 13) // reusing the variables and hence can't use the short hand for assignment := as it requires at least one new variable
	fmt.Printf("Working with a function with Named return values\nRectangle\nArea: %f, Perimeter %f\n\n", area, perimeter)

	area_blank, _ := rectProps(8, 8)
	fmt.Printf("Using Blank Identifier\nRectangle Area: %f\n\n", area_blank)

}
