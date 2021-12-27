package main

import "fmt"

func main(){
	ayam, bebek, kuda := 12, 15 , "Hihihihi"
	fmt.Println(ayam, bebek, kuda)

	bebek, ayam = ayam, bebek
	fmt.Println(ayam, bebek, kuda)
}