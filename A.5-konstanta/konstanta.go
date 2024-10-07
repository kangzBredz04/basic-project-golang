package main

import "fmt"

// Mendefinisikan konstanta menggunakan kata kunci const
const (
	// Konstanta numerik
	PI       = 3.14
	Width    = 10
	Height   = 20
	Username = "Wahyu"
)

func Konstanta() {
	// Menggunakan konstanta dalam perhitungan
	area := Width * Height
	circumference := 2 * PI * Width

	// Mencetak nilai konstanta
	fmt.Println("Konstanta:")
	fmt.Printf("PI: %.2f\n", PI)
	fmt.Printf("Username: %s\n", Username)

	// Mencetak hasil perhitungan
	fmt.Println("\nHasil Perhitungan:")
	fmt.Printf("Luas: %d\n", area)
	fmt.Printf("Keliling: %.2f\n", circumference)
}
