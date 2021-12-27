package main

import "fmt"

func main(){

	var nilai2 [3]int
	nilai2[0] = 0
	nilai2[1] = 10
	nilai2[2] = 20

	var nilai3 = [3]int{1, 2, 3, }
	var nilai4 = [...]int{1, 2, 3, 4, }

	fmt.Println(nilai2)
	fmt.Println(nilai3)
	fmt.Println(nilai4)

	fmt.Println(len(nilai2))
	fmt.Println(len(nilai3))
	fmt.Println(len(nilai4))
}