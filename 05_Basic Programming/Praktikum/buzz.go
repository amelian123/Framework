package main

import "fmt"

func fizz(num int) {
	length := num
	for i := 1; i <= length; i++ {
		if i%3 == 0 && 1%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
func main() {
	fizz(100)
}
