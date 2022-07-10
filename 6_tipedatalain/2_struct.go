package main

import "fmt"

// Membuat Enum
type Jabatan int

const (
	Bos Jabatan = iota
	Sekretaris
	Bendahara
	Bawahan
)

// Membuat Struct
type Orang struct {
	nama string
	umur int
}

type Pasangan struct {
	nama string
	umur int
}

type Pegawai struct {
	Orang // Mengembeddkan, mirip inheritance
	Pasangan
	jabatan Jabatan
	umur int
}

func (o Orang) Halo(){
	fmt.Printf("Halo, %s\n", o.nama)
}

func (o Orang) Haloha(){
	o.Halo()
}

func (p Pegawai) Haloha() {
	fmt.Println("Dari pegawai")
	p.Halo()
}

func main() {
	orang := &Pegawai{}

	// Bisa gini
	orang.jabatan = Bawahan
	// orang.nama = "Fulan" // Ambigous!! Kudu dispesifikan klo anaknya
	orang.Orang.nama = "Fulan"

	// Nyari dulu yg di pegawai, klo gaada baru anaknya
	orang.umur = 20 

	// Bole juga gini
	orang.Orang.nama = "Fulan bin fulan"

	fmt.Println(orang)

	orang.Halo() // Anaknya yg dipanggil karena gaada di parent
	orang.Haloha() // Parent yg dipanggil
	orang.Orang.Haloha() // Spesifik!
}