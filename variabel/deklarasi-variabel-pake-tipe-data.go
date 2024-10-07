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
<<<<<<< HEAD
	fmt.Printf("Hallo nama lengkap saya %s %s\n", firstName, lastName)

	main2()
	MultiVariabel()
	New()
=======
	fmt.Printf("Hallo nama lengkap saya %s %s", firstName, lastName)
>>>>>>> 6beed369288d0e7692753007c690b6222dc13473
}
