package main

import "fmt"

//import "math" - was needed to test math functions here.

func main() {
	var a int = 50
	var b string = "I love Go - may be"
	fmt.Println("var a:", a)
	fmt.Println("var b:", b)
	fmt.Println(`Internally to Go, above values assigned to both variables are constants to Go.`)

	fmt.Println(`
  const cons_a = 55
  cons_a = 89
  gives 'cannot assign to cons_a' as Constants cannot be reassigned.
  `)

	// test assigning value of a constant at compile time
	fmt.Println(`
  var comp_a = math.Sqrt(4)
  const comp_b = math.Sqrt(4)
  gives 'const initializer math.Sqrt(4) is not a constant'
  `)

	fmt.Println("//String Constants")
	//Value enclosed between double quotes, eg: "what's up"
	const hello = "sup world"
	fmt.Printf("type: %T\nvalue %v\n", hello, hello)
	fmt.Println(`
  - String constant belong to type 'untyped'.
  - Above constant 'hello' doesn't have a type.
  - untyped constants have a default type associated with them, which is supplied only if code demands it.
  `)

	fmt.Println(`
  //Creating a typed constant
  //typedhello is a constant of type string
  const typedhello string = "Hello World"
  `)
	const typedhello string = "Hello World"
	fmt.Printf("type: %T\nvalue %v\n", hello, hello)

	// mixing types during assignment isn't allowed
	var defaultname = "batman"         //defaultname type is string type after assignment of constant of string type.
	type mystring string               //create a new type mystring as an alias of string
	var customname mystring = "batman" //create a var customname of type mystring and assign it constant "batman" which is untyped.
	fmt.Println(`
  customname = defaultname
  gives 'cannot use defaultname (type string) as type mystring in assignment'
  because mystring is an alias of string which prevents assignment of defaultname value to customname due to Go's strong typing policy.

   defaultname and customname:
  `, defaultname, customname)

	fmt.Println("\n//Boolean Constants")
	const trueConst = true            // assign untyped constant 'true' to trueConst
	type myBool bool                  //create a new type myBool as an alias to bool type
	var defaultBool = trueConst       //create defaultBool assigning it trueConst; here defaultBool will infer type from trueConst
	var customBool myBool = trueConst //create customBool of type myBool and assign it trueConst
	fmt.Println(`
  defaultBool = customBool //try assigning customBool (type: myBool which is an alias of type bool) to defaultBool (type: bool)
  gives 'cannot use customBool (type myBool) as type bool in assignment'
  `, defaultBool, customBool)

	fmt.Println("\n//Numeric Constants")

	const const_a = 5 // const const_a is untyped.
	var intVar int = const_a
	var int32Var int32 = const_a
	var float64Var float64 = const_a
	var complex64Var complex64 = const_a
	fmt.Println("intVar:", intVar, "\nint32Var:", int32Var, "\nfloat64Var:", float64Var, "\ncomplex64Var:", complex64Var)
	fmt.Println(`
  Syntax of 'const consta_a = 5' is generic and it's possible to assign it to any compatible type.
  `)

	fmt.Println(`
  //Numeric Expressions
  //Numeric constants are free to be mixed and matched in expressions, needing a type only when being assigned to variables or where a type is explicitly demaded.
  `)

	var num_exp_a = 5.9 / 8
	fmt.Printf("num_exp_a type: %T, value: %v\n", num_exp_a, num_exp_a)
	fmt.Println("Expression results in a float and hence float type.")

	fmt.Println(`
  Summary:
  - constants cannot be reassigned.
  - value of constant should be known at compile time.
  `)
}
