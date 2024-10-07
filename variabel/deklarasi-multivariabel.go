package main

import "fmt"

func MultiVariabel() {
	// Deklarasi
	var satu, dua, tiga string
	// Pengisian variabel
	satu, dua, tiga = "satu", "dua", "tiga"

	// Deklarasi dan pengisian variabel
	var empat, lima string = "empat", "lima"

	// Deklarasi dan pengisian variabel yang lebih singkat
	enam, tujuh := "enam", "tujuh"

	// Menampilkan isi variabel satu, dua, tiga
	fmt.Println("Isi nya :", satu, dua, tiga, empat, lima, enam, tujuh)

	// Deklarasi dan pengisian variabel dengan berbeda tipe data
	delapan, lakiLaki, desimal, sapa := 8, true, 2.2, "Hai"

	// 	// Menampilkan isi variabel satu, dua, tiga
	fmt.Println("Isi nya :", delapan, lakiLaki, desimal, sapa)

}
