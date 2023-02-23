package main

import "fmt"

func main() {
	n := 5
	for i := 1; i <= n; i++ {
		for s := i; s < n; s++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print(" *")
		}
		fmt.Println()
	}
}
