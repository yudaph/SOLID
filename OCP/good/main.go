package good

//Definisi ENUM Warna
type Warna int

//Definisi ENUM Ukuran
type Ukuran int

type Barang struct {
	Nama  string
	Warna
	Ukuran
}

type Filter interface {
	IsSatisfied(p *Barang) bool
}

type FilterWarna struct {
	Warna
}

func (spec FilterWarna) IsSatisfied(p *Barang) bool {
	return p.Warna == spec.Warna
}

type FilterUkuran struct {
	Ukuran
}

func (spec FilterUkuran) IsSatisfied(p *Barang) bool {
	return p.Ukuran == spec.Ukuran
}

type AndFilter struct {
	First, Second Filter
}

func (spec AndFilter) IsSatisfied(p *Barang) bool {
	return spec.First.IsSatisfied(p) &&
		spec.Second.IsSatisfied(p)
}

type BetterFilter struct {}

func (f *BetterFilter) Filter(
	products []Barang, filter Filter) []*Barang {
	result := make([]*Barang, 0)
	for i, v := range products {
		if filter.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}