package main

import "fmt"

func main() {
	var sisiatas, sisibawah, tinggi float64
	fmt.Print("input sisi atas:")
	fmt.Scanln(&sisiatas)
	fmt.Print("input sisi bawah:")
	fmt.Scanln(&sisibawah)
	fmt.Print("input tinggi:")
	fmt.Scanln(&tinggi)

	luas := (sisiatas + sisibawah) * tinggi / 2
	fmt.Println("luas trapesium adalah : ", luas)
}
