package main

import (
	"fmt"
	"slp/bad"
	"slp/good"
	"slp/good/utils"
)

func main() {
	//Contoh penggunaan SLP yang tidak sesuai aturan
	fmt.Println("Bad example")
	kbad := bad.Kabinet{}
	kbad.AddEntry("Dokumen 1")
	kbad.AddEntry("Dokumen 2")

	fmt.Println(kbad.String())
	kbad.Save("kabinetBad.txt")

	//Contoh penggunaan SLP sesuai aturan
	fmt.Printf("\n\nGood example\n")
	kgood := good.Kabinet{}
	kgood.AddEntry("Dokumen 3")
	kgood.AddEntry("Dokumen 4")


	fmt.Println(kgood.String())

	p := utils.Persistence{LineSeparator: "\r\n"}
	p.SaveToFile(&kgood, "kabinetGood.txt")
}