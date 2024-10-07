package main

import "fmt"

func main() {
	// Contoh menggunakan if, else if, else
	age := 20

	if age < 18 {
		fmt.Println("Anda masih anak-anak.")
	} else if age >= 18 && age < 60 {
		fmt.Println("Anda adalah orang dewasa.")
	} else {
		fmt.Println("Anda adalah lansia.")
	}

	// Contoh menggunakan switch
	grade := "B"

	switch grade {
	case "A":
		fmt.Println("Nilai Anda sangat baik!")
	case "B":
		fmt.Println("Nilai Anda baik.")
	case "C":
		fmt.Println("Nilai Anda cukup.")
	default:
		fmt.Println("Nilai Anda tidak terdaftar.")
	}

	// Contoh switch dengan banyak kondisi
	day := 4

	switch {
	case day == 1:
		fmt.Println("Hari ini adalah Senin.")
	case day == 2:
		fmt.Println("Hari ini adalah Selasa.")
	case day == 3:
		fmt.Println("Hari ini adalah Rabu.")
	case day == 4:
		fmt.Println("Hari ini adalah Kamis.")
	default:
		fmt.Println("Hari tidak dikenali.")
	}

	// Contoh penggunaan fallthrough
	score := 75

	switch {
	case score >= 90:
		fmt.Println("Nilai A")
	case score >= 80:
		fmt.Println("Nilai B")
		fallthrough // Akan melanjutkan ke case berikutnya
	case score >= 70:
		fmt.Println("Nilai C")
	default:
		fmt.Println("Nilai D")
	}

	// Contoh seleksi kondisi bersarang
	number := 15

	if number%2 == 0 {
		fmt.Println(number, "adalah bilangan genap.")
	} else {
		fmt.Println(number, "adalah bilangan ganjil.")
		if number%3 == 0 {
			fmt.Println(number, "adalah bilangan ganjil yang juga merupakan kelipatan 3.")
		} else {
			fmt.Println(number, "adalah bilangan ganjil yang bukan kelipatan 3.")
		}
	}
}
