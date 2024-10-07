package main

import (
	"fmt"
	"komentar/helpers"
)

func main() {
	// Ini adalah komentar
	// Dibawah ini akan menampilkan hello world
	fmt.Println("Hello World")

	// Komentar tidak akan dieksekusi
	// Ini adalah fungsi multiline dari file komentar-multiline.go di pakacge yang sama
	multiline()

	// Ini adalah fungsi multiline dari file komentar-multiline.go di pakacge helpers
	helpers.Multiline()
}
