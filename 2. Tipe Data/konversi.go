package main

import (
	"fmt"
)

func main(){
	fmt.Println("Melakukan Konversi Tipe DATA")

	var kudaNil = "Kuda Nil"
	var K = kudaNil[3:] // Sekelas C rasa Python :v
	var karakterK = string(67) // Type casting
	var baiknya = fmt.Sprint(67) // Cast yg lebih baik

	println(K)
	println(karakterK)
	println(baiknya)
}