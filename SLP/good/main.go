package good

import (
	"fmt"
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

