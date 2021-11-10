package bad

//Definisi ENUM Warna
type Warna int

//Definisi ENUM Ukuran
type Ukuran int

type Barang struct {
	Nama  string
	Warna
	Ukuran
}

type Filter struct {}


func (f *Filter) FilterWarna(SliceBarang []Barang, warna Warna)[]*Barang {
	result := make([]*Barang, 0)

	for i, v := range SliceBarang {
		if v.Warna == warna {
			result = append(result, &SliceBarang[i])
		}
	}

	return result
}

func (f *Filter) FilterBySize(SliceBarang []Barang, ukuran Ukuran) []*Barang {
	result := make([]*Barang, 0)

	for i, v := range SliceBarang {
		if v.Ukuran == ukuran {
			result = append(result, &SliceBarang[i])
		}
	}

	return result
}

func (f *Filter) FilterBySizeAndColor(products []Barang, ukuran Ukuran, warna Warna)[]*Barang {
	result := make([]*Barang, 0)

	for i, v := range products {
		if v.Warna == warna && v.Ukuran == ukuran {
			result = append(result, &products[i])
		}
	}

	return result
}


