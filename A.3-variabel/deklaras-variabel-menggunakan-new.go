package main

import "fmt"

func New() {
	// membuat pointer ke string
	firstName := new(string)
	// mengisi nilai variabel yang ditunjuk oleh pointer
	*firstName = "Wahyu"
	// menampilkan isi dari variabel yang disimpan kedalam pointer
	fmt.Println(*firstName)
}
