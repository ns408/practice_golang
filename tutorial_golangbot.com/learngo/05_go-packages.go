package main

import "fmt"

func main() {

	fmt.Println(`
  //Notes:
  //Go version >= 1.13 uses Go modules.

  //Packages:
  //- are used to organise Go source code for better reusability and readability.
  //- are a collection of Go source files.
  //- provide code compartmentalisation making it easy to maintain projects.

  //main function and main package:
  //Every Go application must contain the main function and main function is the entry point for execution.
  //main function resides in the main package.

  //'package packagename' specifies that a file belongs to package packagename.
  //eg: 'package main' in main.go

  //Go Module:
  //- is a collection of Go packages.
  //- import path for packages is derived from the module name
  //- $module/go.mod contains the packages and their versions. It get created on module creation.

	//Go Packages:
	//- function name in a package starts with Caps. Eg: Calculate, Println

	//from within the module directory, use either of these
	go install
	go install .
	go install learnpackage

	//Exported Names
	'mod_learnpackage/simpleinterest' package has a function 'Calculate' instead of 'calculate'.
	- any variable with a capital letter are **Exported names** in go. Only then can these be accessed from outside the package.

	//changing the function to 'calculate' and then running 'mod_learnpackage/main.go' gives the following error:
	./main.go:13:8: cannot refer to unexported name simpleinterest.calculate
	./main.go:13:8: undefined: simpleinterest.calculate

	//init function
	- each package can contain an 'init' function.
	- it doesn't have a return type and parameters.
	- it can't be called explicitly.
	- it is called automatically when the package is initialised.
	- it can be used to perform initialisation tasks and to verify the corectness of the program before execution starts.
		- package level variables are initialised first.
		- init function is called next. A package can have mulitple init functions in a single or multiple files. They are called in the order they are presented to the compiler.
	- if a package imports other packages, the imported packages are initialised first.
	- a package is initialised only once even if imported from multiple packages.

	syntax:
	func init() {
	}

	//blank identifier
	- illegal to import a package and not use it. This avoids bloating of unused packages which increases compile time.

	Example of use:
	import "module/package"
	var _ = "package.FunctionName

  `)

	fmt.Println(`
  //Steps for module creation:
		mkdir mod_learnpackage
		vim mod_learnpackage/main.go // populate this with the content
		go fmt mod_learnpackage/main.go // format the file
		cd mod_learnpackage; go mod init mod_learnpackage // creates 'go.mod'
	//mod_learnpackage will be the base path to import any packages inside the module.

	//Create the simpleinterest custom package
	//source files belonging to a package should be placed in separate folder.
  //Go convention is to name the directory same name as the package.
	mkdir mod_learnpackage/simpleinterest
	vim mod_learnpackage/simpleinterest/simpleinterest.go // populate the file
  //All files in mod_learnpackage/simpleinterest should start with 'package simpleinterest' as they are all part of the same package.

	//Test the module/package
	cd mod_learning/; go run main.go

	More play with Named return values. Also mucked around with GOPATH.
	Seemed that only two ways of utilising the module/package was to either:
	- mkdir src && mv mod_learnpackage src && export GOPATH=$(PWD)
	- cp -av mod_learnpackage "$(go env PATH)/src/"
  `)

}
