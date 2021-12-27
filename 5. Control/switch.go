package main

import "fmt"

func main(){
	kumbang := "HAHA"

	/* Operator pembandingnya adalah ==. Bisa menerima tipe data apa aja */
	switch kumbang {
		case "HEHE":
			fmt.Println("Hohoho...")
			break // Tidak wajib, udah otomatis
		case "HAHA":
			fmt.Println("Hehehehe...")
		case "HOHO":
			fmt.Println("Hahaha")
		default:
			fmt.Println("Ini default")
	}

	hantu := 666

	switch nilai := hantu + 10; nilai {
		case 676:
			fmt.Println("Kudanil")
		default:
			fmt.Println("Angsa")
	}

	/* Bahkan boleh gini */
	
	switch{
		case hantu == 666:
			fmt.Println("Kikikikik")
	}
}