package main

import (
	"fmt"
	"math"
)

// 1. Ini cara membuat interface
type Hitung2d interface {
	Luas() float64
	Keliling() float64
}

type Hitung3d interface {
	Volume() float64
	LuasPermukaan() float64
}

// 2. Bisa juga diembedd jadi

type Hitung interface {
	Hitung2d
	Hitung3d
}

type Tabung struct {
	jariJari float64
	tinggi float64
}

func (t Tabung) Luas() float64 {
	return t.jariJari * t.jariJari * math.Acos(-1);
}

func (t Tabung) Keliling() float64 {
	return 2 * t.jariJari * math.Acos(-1);
}

func (t Tabung) Volume() float64 {
	return t.Luas() * t.tinggi;
}

func (t Tabung) LuasPermukaan() float64 {
	return t.Keliling() * t.tinggi + 2 * t.Luas() 
}

func main() {
	var kudanil Hitung = Tabung{jariJari: 7, tinggi: 4.0}
	var kiki interface{} = kudanil // Bisa apa aja
	
	fmt.Println(kudanil.Volume())
	fmt.Println(kudanil.LuasPermukaan())
	fmt.Println(kiki)
}