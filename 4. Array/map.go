package main

import "fmt"

func main(){
	makanan := map[string]int{
		"nasi goreng": 10,
		"udon": 5,
	}

	makanan["enak"] = 100

	fmt.Println(makanan)
	fmt.Println(makanan["udon"])
	fmt.Println(makanan["gaada loh"]) // Diisi 0
	fmt.Println(makanan["enak"])
	fmt.Println(len(makanan)) // 3
	fmt.Println()
	fungsiPembantu()
}

func fungsiPembantu(){
	makananBaru := make(map[string]int)
	makananBaru = map[string]int{} // Sah" aja
	
	makananBaru["wakakak"] = 10
	makananBaru["hupla"] = 10
	fmt.Println(makananBaru)
	
	delete(makananBaru, "hupla")
	fmt.Println(makananBaru)
}