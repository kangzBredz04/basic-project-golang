package main

import "fmt"

func main() {
	// Deklarasi dan inisialisasi map
	ages := make(map[string]int)

	// Menambahkan data ke dalam map
	ages["Wahyu"] = 25
	ages["Budi"] = 30
	ages["Siti"] = 28

	// Mencetak isi map
	fmt.Println("Ages:", ages) // Output: Ages: map[Budi:30 Siti:28 Wahyu:25]

	// Mengakses nilai dari map
	wahyuAge := ages["Wahyu"]
	fmt.Println("Wahyu's age:", wahyuAge) // Output: Wahyu's age: 25

	// Mengubah nilai dalam map
	ages["Budi"] = 31
	fmt.Println("Updated Ages:", ages) // Output: Updated Ages: map[Budi:31 Siti:28 Wahyu:25]

	// Menghapus elemen dari map
	delete(ages, "Siti")
	fmt.Println("Ages after deletion:", ages) // Output: Ages after deletion: map[Budi:31 Wahyu:25]

	// Mengecek apakah kunci ada di dalam map
	age, exists := ages["Siti"]
	if exists {
		fmt.Println("Siti's age:", age)
	} else {
		fmt.Println("Siti not found in ages map") // Output: Siti not found in ages map
	}

	// Menggunakan loop untuk mencetak semua elemen dalam map
	fmt.Println("All ages:")
	for name, age := range ages {
		fmt.Printf("%s: %d\n", name, age)
	}
}
