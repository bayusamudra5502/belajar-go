package data

import (
	"fmt"
	"reflect"
)

type Orang struct {
	nama string
	Kelas int
	skor int
}

func (o *Orang) Jadiin100() {
	o.skor = 100
}

func NewOrang(nama string) Orang {
	return Orang{nama: nama, Kelas: 1, skor: 0}
}

func (o Orang) Print() {
	println("Nama : ", o.nama)
	println("Kelas : ", o.Kelas)
	println("Skor : ", o.skor)
}

func NamaNam(source interface{}, nama string){
	fmt.Println(reflect.ValueOf(source).CanAddr())

	// reflectOrang := reflect.ValueOf(source).Addr() // Panikan diamah klo gabisa 
	reflectOrang := reflect.ValueOf(source).Elem() // Klo ini bakal jadi nal nil nul

	fieldNama := reflectOrang.FieldByName("nama")

	fmt.Println(">",fieldNama.CanSet())
	fmt.Println(">", fieldNama.String())
}
