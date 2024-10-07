package main

import "fmt"

func main() {
	// Operator Aritmatika
	a := 10
	b := 3

	// Penjumlahan
	sum := a + b
	fmt.Printf("Penjumlahan: %d + %d = %d\n", a, b, sum)

	// Pengurangan
	sub := a - b
	fmt.Printf("Pengurangan: %d - %d = %d\n", a, b, sub)

	// Perkalian
	mul := a * b
	fmt.Printf("Perkalian: %d * %d = %d\n", a, b, mul)

	// Pembagian
	div := a / b
	fmt.Printf("Pembagian: %d / %d = %d\n", a, b, div)

	// Sisa bagi (modulus)
	mod := a % b
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, mod)

	// Operator Perbandingan
	fmt.Printf("Apakah %d > %d? %t\n", a, b, a > b)
	fmt.Printf("Apakah %d < %d? %t\n", a, b, a < b)
	fmt.Printf("Apakah %d == %d? %t\n", a, b, a == b)
	fmt.Printf("Apakah %d != %d? %t\n", a, b, a != b)
	fmt.Printf("Apakah %d >= %d? %t\n", a, b, a >= b)
	fmt.Printf("Apakah %d <= %d? %t\n", a, b, a <= b)

	// Operator Logika
	x := true
	y := false

	fmt.Printf("x AND y: %t\n", x && y) // AND
	fmt.Printf("x OR y: %t\n", x || y)  // OR
	fmt.Printf("NOT x: %t\n", !x)       // NOT
}
