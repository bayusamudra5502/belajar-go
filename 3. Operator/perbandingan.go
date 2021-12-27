package main

import "fmt"

func main(){
	a := 12
	var b uint8 = 3

	// Harus Sama Tipedatanya!
	// Go gaboleh tebak"an tipe data ya wkwkwk
	fmt.Println(a > int(b))
}