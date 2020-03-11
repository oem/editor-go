package piecetable

import "io/ioutil"

type pieceTable struct {
	original string
	add      string
	pieces   []piece
}

type piece struct {
	moo int
}

func New(buf string) *pieceTable {
	return &pieceTable{original: buf}
}

func NewFromFile(filename string) (*pieceTable, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &pieceTable{original: string(f)}, nil
}
