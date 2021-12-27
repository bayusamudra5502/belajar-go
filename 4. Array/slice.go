package main

import "fmt"

func main(){
	 var nilai = [...]int{
		1,
		2,
		3,
		4,
	}

	// Awas ini beda:
	var iniSlice = []int{
		1,
		2,
		3,
		4,
	}

	// Cara lain membuat Slice
	const CAPACITY = 10
	const LENGTH = 5
	sliceCaraLain := make([]int, LENGTH, CAPACITY)
	sliceCaraLain[0] = 12
	sliceCaraLain[1] = 12

	fmt.Println("Slice Cara lain",sliceCaraLain)

	fmt.Println(iniSlice)

	/* Slice terhubung sama array utama, artinya
		klo slice diubah atau arraynya diubah akan
		mengubah semuanya (array dan slice) */

	var slice1 = nilai[2:]
	fmt.Println(slice1)
	
	slice1[0] = 20
	fmt.Println(slice1)
	fmt.Println(nilai)
	
	nilai[3] = 50
	fmt.Println(slice1)
	fmt.Println(nilai)
	fmt.Println()
	fungsiSlice()

	fmt.Println()
	copySlicekuh()
}

func fungsiSlice(){
	kucing := []int{
		1,
		2,
		3,
		4,
		5,
	}

	kuda := [...]int{
		2,
		3,
		4,
		5,
		6,
		7,
	}

	fmt.Println(len(kucing))
	fmt.Println(cap(kucing))

	fmt.Println(len(kuda[:5]))
	fmt.Println(cap(kuda[:5]))

	// Fungsi tambahan
	// Karena kucing udah penuh dibuat array baru
	// dan pengubahan tidak mengubah arraynya
	kucingBaru := append(kucing, 10)
	fmt.Println(kucingBaru)
	fmt.Println(kucing)

	// Kasus gini bakal ngereplace nilai array kuda
	fmt.Println(kuda)
	var belalang = append(kuda[:2], 15)
	fmt.Println(belalang)
	fmt.Println(kuda)
}

func copySlicekuh(){
	kudanil := []int{
		2,
		4,
		6,
		8,
		10,
	}

	kucing := make([]int, len(kudanil), 10)
	kumbang := make([]int, 2, 10)
	kuda := make([]int, 8, 10)

	copy(kucing, kudanil)
	copy(kumbang, kudanil)
	copy(kuda, kudanil)

	fmt.Println(kucing)
	fmt.Println(kumbang)
	fmt.Println(kuda)
}