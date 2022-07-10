package main

import (
	"fmt"
	"strconv"
)

// Klo berkaitan ngubah string jadi angka"an, coba
// cek strconv

func main(){
	satu := "120"

	angka, err := strconv.Atoi(satu)
	
	if err != nil{
		panic(err)
	}

	fmt.Printf("Angkanya: %d\n", angka + 10)
}