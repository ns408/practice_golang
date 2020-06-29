
package main
import "fmt"
//import "math" - was needed to test math functions here.

func main() {
  var a int = 50
  var b string = "I love Go - may be"
  fmt.Println("var a:", a )
  fmt.Println("var b:", b )
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


  fmt.Println(`
  Summary:
  - constants cannot be reassigned.
  - value of constant should be known at compile time.
  `)
}
