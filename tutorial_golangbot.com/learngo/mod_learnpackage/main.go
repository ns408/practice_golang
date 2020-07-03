package main

/* Order of initialisation
- imported packages first and their init method is called.
- package level vars
- init function is called in main package
- main function
*/

import (
	"mod_learnpackage/simpleinterest" // importing the package into this module
)

var _ = simpleinterest.Calculate // mutes the error - but should be used with caution

func main() {
}

/* Importing the package above and not using it gives the following error:
# command-line-arguments
./main.go:11:2: imported and not used: "mod_learnpackage/simpleinterest"
*/
