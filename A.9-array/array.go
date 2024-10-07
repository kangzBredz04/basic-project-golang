package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"
	names[3] = "d"
	// Isi dari array
	fmt.Println(names[0], names[1], names[2], names[3])

	// Panjang array
	fmt.Println(len(names))

	var fruits [4]string
	// cara horizontal
	fruits = [4]string{"apple", "grape", "banana", "melon"}
	// cara vertikal
	fruits = [4]string{
		"apple",
		"grape",
		"banana",
		"melon",
	}

	// Isi dari array fruits
	fmt.Println(fruits)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var numbers = [...]int{2, 3, 2, 4, 3}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))

	// Membuat array dua dimensi 3x3
	var matrix [3][3]int

	// Mengisi array dengan nilai
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i + j // Mengisi dengan penjumlahan indeks
		}
	}

	// Mencetak array dua dimensi
	fmt.Println("Matrix:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println() // Pindah ke baris baru setelah mencetak satu baris
	}

	// Contoh akses elemen tertentu
	fmt.Printf("Element at (1, 2): %d\n", matrix[1][2]) // Output elemen di baris 1 kolom 2

	// Perulangan Elemen Array Menggunakan Keyword for - range
	// Deklarasi dan inisialisasi array
	months := [12]string{
		"Januari", "Februari", "Maret", "April",
		"Mei", "Juni", "Juli", "Agustus",
		"September", "Oktober", "November", "Desember",
	}

	// Menggunakan for-range untuk perulangan elemen array
	fmt.Println("Daftar Bulan:")
	for index, month := range months {
		fmt.Printf("Bulan ke-%d: %s\n", index+1, month)
	}

	// Alokasi Elemen Array Menggunakan Keyword make
	var fruits2 = make([]string, 2)
	fruits2[0] = "apple"
	fruits2[1] = "manggo"
	fmt.Println(fruits2)
}
