package main

import "fmt"

type fungsiMatematika func(int) int

func main(){
	/* Harus ditangkep semua klo ga error */
	a, b, c := Belalang()

	/* Cara skip pakai _ */
	a, _, c = Belalang()

	fmt.Println(a, b, c)
	fmt.Println(f(10))
	fmt.Println(gege(20))
	fmt.Println(variadicFunction(1,2,3,4,5,6))

	/* Rada lain dari js wkwkwk, disimpennya di belakang :v 
		 untuk destructuring */
	fmt.Println(variadicFunction([]int{1,2,3,4,5,6}...))

	/* Kek C dan kawan"nya bisa gini kita */
	kumbang := f
	fmt.Println(kumbang(25))

	/* Bole malah gini */
	kucing := func(a int) (int, int) {
		return a * 2, a + 2
	}

	fmt.Println(kucing(12))
}

func Belalang() (int, int, string){
	return 1, 2, "Baaa"
}

func f(bacang int) int{
	return bacang * 2
}

/* Named return value */
func gege(empat byte) (satu float32, dua, tiga int) {
	println(empat)

	/* Sudah dideklarasikan di atas. Tinggal inisilaisasi aja */
	satu = 1.9
	dua = 15
	tiga = 10

	return satu, dua, tiga
}

func variadicFunction(angka ...int) float64 {
	sum := 0

	for _, number := range angka {
		sum += number
	}

	return float64(sum) / float64(len(angka))
}

func reducer(angka []int, fungsi fungsiMatematika){
	
}