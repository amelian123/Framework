package main

import "fmt"

func main() {
	fmt.Print("Input Bilangan :")
	var num int
	fmt.Scanf("%d", &num)

	if (num % 2) == 0 {
		fmt.Printf("%d adalah bilangan genap", num)
	} else {
		fmt.Printf("%d adalah bilangan ganjil", num)
	}
}
