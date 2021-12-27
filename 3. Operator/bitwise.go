package main

import "fmt"

func main(){
	var a = 10
	var b = 0b10111
	var c = 0xBACA

	fmt.Println(a & b)
	fmt.Println(b | c)
	fmt.Println(a ^ b) // XOR
	fmt.Println(^a) // Negasi a, klo di C ma ~a
	fmt.Println(a << 2)
	fmt.Println(a >> 2)

}