package main

import (
	"fmt"
	"mod_learnpackage/simpleinterest"
)

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
  `)

	fmt.Println("Simple interest calculation")
	p := 1000.0
	r := 7.0
	t := 1.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
