package main

/* Order of initialisation
- imported packages first and their init method is called.
- package level vars
- init function is called in main package
- main function
*/

import (
	"fmt"
	"log"
	"mod_learnpackage/simpleinterest" // importing the package into this module
)

var p, r, t = 5000.0, 10.0, 1.0 // variables moved here from the main() function
/*var p, r, t = -5000.0, 10.0, 1.0 // variables moved here from the main() function
produces:
Simpeinterest package initialised
Main package initiliased
2020/07/03 22:29:26 Principal < 0
exit status 1
*/

/*
* Use init function to check that the principal, rate and duration are greater than zero.
 */
func init() {
	println("Main package initiliased")
	if p < 0 {
		log.Fatal("Principal < 0")
	}
	if r < 0 {
		log.Fatal("rate < 0")
	}
	if t < 0 {
		log.Fatal("duration < 0")
	}

}

func main() {
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
