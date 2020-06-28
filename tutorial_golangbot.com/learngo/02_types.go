
//Notes:
//Basic data types:
//  bool
//  numeric
//  - int8, int16, int32, int64, int
//  - uint8, uint16, uint32, unit64, uint
//  - float32, float64
//  - complex64, complex128
//  - byte
//  - rune
//  string

package main
import "fmt"
import "unsafe"
//Importing for `Sizeof` function - returns in bytes the var size
//use unsafe package with care - portability issues

//Another way:
//import (
//  "fmt"
//  "unsafe"
//)

func main() {
  a := true
  b := false
  fmt.Println("a:", a, "b:", b)
  c := a && b
  fmt.Println("c:", c)
  d := a || b
  fmt.Println("d:", d)

  var var_a int = 89
  var_b := 95 // inferred type int
  fmt.Println("value of var_a is", var_a, "and var_b is", var_b)
  //verify the size of int on a 64 bit (8 bytes) system
  fmt.Printf("var_a type: %T, var_a size: %d\n", var_a, unsafe.Sizeof(var_a))
  fmt.Printf("var_b type: %T, var_b size: %d\n", var_b, unsafe.Sizeof(var_b))

  // Floating point types
  var_a1, var_b1 := 5.5, 10.55
  fmt.Printf("var_a1 type: %T, var_b1 type: %T\n", var_a1, var_b1)
  sum := var_a1 + var_b1
  diff := var_a1 - var_b1
  fmt.Println("sum", sum, "diff", diff)

  // Complex types
  // builtin function `complex` is used to construct a complex number (real and imaginary parts)
  // Definition: func complex(r, i Floattype) ComplexType
  // Both real and imaginary parts should be of the same type - either float32 or float64
  // Where both real and imaginary parts are float64 type, complex value returned is of type complex128

  complex_a1 := complex(10, 20)
  complex_b1 := 7 + 13i
  cadd := complex_a1 + complex_b1
  fmt.Println("sum:", cadd)
  cmul := complex_a1 * complex_b1
  fmt.Println("product:", cmul)

  //String type
  // Collections of bytes in Go.
  fname := "Batman"
  lname := "Joker"
  fullname := fname + " " + lname
  fmt.Println("The names are", fullname)

  //Type conversion
  // Explicit typing is needed; no automatic type promotion or conversion.

  //i := 5
  //j := 10.5
  //fmt.Println(i + j)
  //gives `invalid operation: i + j (mismatched types int and float64)`, though it's legal in C language.

  i := 5
  j := 10.5
  //var k float64 = i
  //gives `cannot use i (type int) as type float64 in assignment`
  var k float32 = float32(i)

  fmt.Println("Explicit type conversion for operation:", i + int(j))
  fmt.Println("Explicit type conversion needed for value assignment:", k)
}
