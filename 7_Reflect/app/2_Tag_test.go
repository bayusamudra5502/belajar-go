package app

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bayusamudra5502/belajar-go/reflect/data"
)

func TestTag(t *testing.T) {
	mhs1 := data.Mahasiswa{Nama: "Bayu", Nim: 13520128, NomorPendaftaran: 12}
	strc := reflect.TypeOf(mhs1)
	vals := reflect.ValueOf(mhs1)

	for i := 0; i < strc.NumField(); i++ {
		field := strc.Field(i)
		fmt.Printf("[%s] %s=", field.Type.Kind(),field.Tag.Get("field"))
		fmt.Println(vals.Field(i).Interface())
	}
}