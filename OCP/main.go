package main

import (
	"fmt"
	"ocp/bad"
	"ocp/good"
)

//Definisi ENUM Warna
const (
	merah bad.Warna = iota
	kuning
	biru
)
const (
	hijau good.Warna = iota
	ungu
	orange
)

//Definisi ENUM Ukuran
const (
	kecil bad.Ukuran = iota
	sedang
	besar
)

func main() {
	//Bad Example
	fmt.Println("BAD EXAMPLE")
	a := bad.Barang{
		Nama: "Barang A",
		Ukuran: besar,
		Warna: merah,
	}
	b := bad.Barang{
		Nama: "Barang B",
		Ukuran: kecil,
		Warna: biru,
	}
	sliceBarang := []bad.Barang{a,b}
	bfilter := bad.Filter{}
	for _, v := range bfilter.FilterBySize(sliceBarang, kecil) {
		fmt.Printf("%s ukuran kecil\n", v.Nama)
	}
	for _, v := range bfilter.FilterWarna(sliceBarang, merah) {
		fmt.Printf("%s warna merah\n", v.Nama)
	}
	for _, v := range bfilter.FilterBySizeAndColor(sliceBarang, besar, merah) {
		fmt.Printf("%s ukuran besar dan warna merah\n", v.Nama)
	}


	//Good Example
	fmt.Println("GOOD EXAMPLE")
	c := good.Barang{
		Nama: "Barang C",
		Ukuran: good.Ukuran(kecil),
		Warna: ungu,
	}
	d := good.Barang{
		Nama: "Barang D",
		Ukuran: good.Ukuran(kecil),
		Warna: hijau,
	}
	sliceBarang2 := []good.Barang{c,d}
	hijauFilter := good.FilterWarna{hijau}
	bf := good.BetterFilter{}
	for _, v := range bf.Filter(sliceBarang2, hijauFilter) {
		fmt.Printf(" - %s is hijau\n", v.Nama)
	}
	kecilFilter := good.FilterUkuran{good.Ukuran(kecil)}
	for _, v := range bf.Filter(sliceBarang2, kecilFilter) {
		fmt.Printf(" - %s is kecil\n", v.Nama)
	}

	andFilter := good.AndFilter{hijauFilter, kecilFilter}
	for _, v := range bf.Filter(sliceBarang2, andFilter) {
		fmt.Printf(" - %s is hijau dan kecil\n", v.Nama)
	}

}