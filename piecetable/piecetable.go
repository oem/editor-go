package piecetable

import "io/ioutil"

type table struct {
	original string
	add      string
	pieces   []piece
}

type piece struct {
	offset int
	length int
	added  bool
}

func New(buf string) *table {
	table := &table{original: buf}
	table.pieces = append(table.pieces, piece{offset: 0, length: len(table.original), added: false})
	return table
}

func NewFromFile(filename string) (*table, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(string(f)), nil
}

func (pt *table) Delete(start, end int) {
	before := piece{offset: 0, length: start + 1, added: false}
	after := piece{offset: end + 1, length: len(pt.original) - (end + 1), added: false}
	pt.pieces = []piece{before, after}
}
