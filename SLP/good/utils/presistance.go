package utils

import (
	"io/ioutil"
	"slp/good"
	"strings"
)

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *good.Kabinet, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.Dokumen, p.LineSeparator)), 0644)
}