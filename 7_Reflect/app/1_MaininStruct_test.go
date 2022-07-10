package app

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/bayusamudra5502/belajar-go/reflect/data"
)

func TestOrang(t *testing.T) {
	orang := data.NewOrang("Bayu")
	orang.Print()

	reflectOrang := reflect.ValueOf(orang) // AMBIL VALUE!!
	field := reflectOrang.FieldByName("Kelas")
	fieldSkor := reflectOrang.FieldByName("skor")
	// field.SetInt(3)

	fmt.Println()
	fmt.Println(field.CanInt()) // Klo bisa ngakses dia true
	fmt.Println(field.CanInterface())

	fmt.Println(field.Int()) // Dapetin Value Integer
	fmt.Println(field.Interface()) // Dapetin Value Integer
	
	fmt.Println()
	fmt.Println(fieldSkor.CanInt()) // Klo bisa ngakses dia true
	fmt.Println(fieldSkor.CanInterface())

	fmt.Println(fieldSkor.Int())

	fmt.Println()
	data.NamaNam(&orang, "Baybay")	
	
	orang.Print()
	
	// Utilitass lain
	
	fmt.Println()
	fmt.Println(fieldSkor.Type())
	fmt.Println(fieldSkor.Kind() == reflect.Int) // Biar bisa dicek
	
	// Bongkarin isi struct
	indi := reflect.Indirect(reflectOrang)

	fmt.Println(indi.Type().NumField())
	fmt.Println(indi.Type().NumMethod())

	// Ato klo mo bongkarin tapi langsung, bisa gini juga
	haha := reflect.TypeOf(orang)
	fmt.Println(haha.Field(0).Name)
	fmt.Println(haha.Method(0).Name)
}

func TestMethod(t *testing.T) {
	orang := data.NewOrang("Bayu")
	reflectOrang := reflect.ValueOf(&orang) // *Orang bakal kenain Jadiin 100
	reflectOrang2 := reflect.ValueOf(orang) 

	val := reflectOrang.Method(0) // Ambil fungsi publik ke-0 (pertama)

	fmt.Println(val.Type())
	fmt.Println(runtime.FuncForPC(val.Pointer()).Name()) // Ngambil nama fungsi
	fmt.Println(runtime.FuncForPC(val.Pointer()).Entry()) // Ngambil alamatnya dimana?

	val.Call([]reflect.Value{});

	var fungsiPrint = reflectOrang2.MethodByName("Print") // Dapetin pake nama
	fungsiPrint.Call([]reflect.Value{})
}
