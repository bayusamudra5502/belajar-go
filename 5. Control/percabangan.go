package main

import "fmt"

func main(){
	umur := 14
	
	/* Ada 2 cara bos */
	if (umur >= 20) {
		fmt.Println("Yay Lebih dari sama dengan 20")
	}

	if umur >= 15 {
		fmt.Println("Yay Lebih dari sama dengan 20")
	} else if umur > 10 {
		fmt.Println("OKee kamu antara 10 s.d 15 tahun")
	} else {
		fmt.Println("Dibawah umur :(")
	}

	/* Bentuk dibawah cuma boleh gitu, gabisa nambah lebih dari 1 
		short statment */
	if kumbang := umur + 5; kumbang > 0 {
		fmt.Println("Haha!!")
	}

	/* Nilai yg diisialisasikan itu punyanya blok lokal yah */
	if umur := 11; umur > 10 {
		fmt.Println("Kujang", umur)
	}

	fmt.Println("Kujang", umur)
}