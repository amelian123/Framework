package main

import "fmt"

func main() {

	var factorsnum, i int

	fmt.Print("Input angka : ")
	fmt.Scanln(&factorsnum)

	fmt.Println("Faktor dari ", factorsnum, " adalah ")
	for i = 1; i <= factorsnum; i++ {
		if factorsnum%i == 0 {
			fmt.Println(i)
		}
	}
}
