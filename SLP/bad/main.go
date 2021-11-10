package bad

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var jumlahDokumen = 0
type Kabinet struct {
	Dokumen []string
}

func (k *Kabinet) String() string {
	return strings.Join(k.Dokumen, "\n")
}

func (k *Kabinet) AddEntry(nama string) int {
	jumlahDokumen++
	input := fmt.Sprintf("%d: %s", jumlahDokumen, nama)
	k.Dokumen = append(k.Dokumen, input)
	return jumlahDokumen
}

func (k *Kabinet) RemoveEntry(index int) {
	if index < jumlahDokumen && index > 0{
		jumlahDokumen--
		k.Dokumen = append(k.Dokumen[:index], k.Dokumen[index+1:]...)
	}
}

//Menyalahi aturan SLP
func (j *Kabinet) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.Dokumen, "\r\n")), 0644)
}


