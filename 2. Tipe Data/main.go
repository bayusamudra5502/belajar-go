package main

import "fmt"

func main(){
	fmt.Printf("Halo, %s\n", "Bayu")
	fmt.Println("Belalang = ", 123)
	fmt.Printf("%c\n","Kuda Lumping"[2])
	fmt.Println(len("Kucing"))
	variable()
	konstanta()
}

func variable(){
	fmt.Println("\nKita bervariabelria")

	/* 4 Cara mendeklarasikan Variabel

		Variabel harus dipake klo enggak ngadat
	*/
	// Cara 1:
	var kucing int

	// Cara 2:
	var kumbang = 50

	// Cara 3:
	belalang := "Tanpa Var bos!"

	// Cara 4:
	var (
		kelinci = "12"
		kumbangLaut = 12
	)

	/* Hanya bisa mengassign dengan tipe data baru */
	/* Berpikir seperti C */
	kucing = 20
	kumbang = 10

	fmt.Println("kucing =", kucing)
	fmt.Println("kumbang =", kumbang)
	fmt.Println(belalang)
	fmt.Println(kelinci)
	fmt.Println(kumbangLaut)
}

func konstanta(){
	/* Konstanta wajib kedah harus langsung diisi dengan valuenya
	
		Konstanta ga dipake gaakan dikomplain sama go-nya.
	*/
	const kuda string = "Hihihihi"
	const belalang = "Krik Krik"

	const (
		kumbang string = "Kumbang Laut"
		kudaNil = "Hihoh"
	)

	fmt.Println(kuda)
	// fmt.Println(belalang) // Bukti
}