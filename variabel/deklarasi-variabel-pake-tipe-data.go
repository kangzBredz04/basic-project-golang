package main

import "fmt"

func main() {
	// Deklarasi sekaligus pengisian data
	var firstName string = "Wahyu"

	// Deklarasi namun pengisian data dipisah
	var lastName string
	lastName = "Jebredz"

	// Ini menggunakan fungsi fmt.Println
	fmt.Println("Hallo nama lengkap saya", firstName, lastName)

	// Ini menggunakan fungsi fmt.Printf
	fmt.Printf("Hallo nama lengkap saya %s %s", firstName, lastName)
}
