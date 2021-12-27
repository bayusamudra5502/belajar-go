package main

import "fmt"

/* OPerator matematika dan bool di go sama kek di C */

func main(){
	nilaiA := 10
	nilaiB := 3
	var nilaiC int8 = 2

	// Tiap operand harus punya tipe data yang sama
	// Klo enggak ngomel!

	var hasil = nilaiA / nilaiB // Div
	var hasil2 = nilaiA / int(nilaiC)
	var hasilFloat float64 = float64(nilaiA) / float64(nilaiB)

	fmt.Println(hasil, hasil2, hasilFloat)
}