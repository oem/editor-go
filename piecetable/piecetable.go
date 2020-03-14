package piecetable

import "io/ioutil"

type pieceTable struct {
	original string
	add      string
	pieces   []piece
}

type piece struct {
	offset int
	length int
	added  bool
}

func New(buf string) *pieceTable {
	table := &pieceTable{original: buf}
	table.pieces = append(table.pieces, piece{offset: 0, length: len(table.original), added: false})
	return table
}

func NewFromFile(filename string) (*pieceTable, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(string(f)), nil
}
