package main

import "fmt"

func main() {
	// 1. Perulangan menggunakan for
	fmt.Println("Perulangan dengan for:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// 2. Perulangan dengan while-style for
	fmt.Println("\nPerulangan dengan while-style for:")
	j := 0
	for j < 5 {
		fmt.Printf("Iteration %d\n", j)
		j++
	}

	// 3. Perulangan menggunakan for dengan slice
	fmt.Println("\nPerulangan dengan slice:")
	fruits := []string{"Apple", "Banana", "Cherry", "Date"}
	for index, fruit := range fruits {
		fmt.Printf("Fruit at index %d: %s\n", index, fruit)
	}

	// 4. Perulangan menggunakan for dengan map
	fmt.Println("\nPerulangan dengan map:")
	ages := map[string]int{
		"Wahyu": 25,
		"Budi":  30,
		"Siti":  28,
		"Joko":  35,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	// 5. Menghentikan perulangan menggunakan break
	fmt.Println("\nMenghentikan perulangan dengan break:")
	for k := 0; k < 10; k++ {
		if k == 5 {
			fmt.Println("Reached 5, breaking the loop.")
			break
		}
		fmt.Println(k)
	}

	// 6. Melanjutkan ke iterasi berikutnya dengan continue
	fmt.Println("\nMelanjutkan dengan continue:")
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			fmt.Printf("%d is even, skipping...\n", m)
			continue
		}
		fmt.Printf("%d is odd\n", m)
	}
}
