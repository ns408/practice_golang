package main

import "fmt"

func main() {

	fmt.Println(`
  // for loop syntax
  for initialisation; condition; post {
  }

  // break

  // continue

  // Nested for loops

  // Labels
	// Can be used to break from an outer loop

	// infinite loop
	for {
	}
  `)

	for i := 1; i <= 10; i++ {
		//fmt.Printf(" %d", i)
		//fmt.Println(i)
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\n")

	// break example
	for i := 1; i <= 10; i++ {
		if i == 6 {
			break //loop is terminated if condition met
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n'break' example ends here")

	// continue example
	// print odd numbers from 1 to 10
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println("\n'continue' example ends here\n")

	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Without 'Labels'
	// simple example
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}
	fmt.Println()

	// Without 'Labels'
	// if i == j, break loop (inner)
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
	fmt.Println()

	// 'Labels'
	// if i == j, break outer loop
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d, j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
	fmt.Println()

	// 'Labels'
	// 3 loops: i, j, k
	// if i == k, break outer loop
outer1:
	for i := 0; i < 3; i++ {
	second:
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				fmt.Printf("i = %d, j = %d, k = %d\n", i, j, k)
				if i == k {
					break second
				} else if i > 10 {
					break outer1
				}
			}
		}
	}
	fmt.Println()

	// initialisation and post ommitted and only condition present
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println("\n")

	// declare and operate on multiple variables
	// let's multiply
	//for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
	fmt.Println()

	// infinite loop
	//for {
	//fmt.Printf(".")
	//}

}
