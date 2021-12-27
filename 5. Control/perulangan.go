package main

import "fmt"

func main(){
	counter := 1
	
	for counter < 20 {
		fmt.Println("cnt ke-", counter)
		counter++ // ++ cuma bole di belakang yah
	}

	/* Counter sebagai variabel lokal di blok kurawal yah */
	for counter := 1; counter < 10; counter++ {
		if counter == 5 {
			continue;
		} else if counter > 8 {
			break;
		}

		fmt.Println(counter)
	}

	fmt.Println(counter)

	/* Kekuatan Ular :v */
	slicekuh := []string { "Ayam", "Bebek", "Domba"}

	/* Destructuring */
	for idx, data := range slicekuh {
		fmt.Println(idx, data)
	}
}