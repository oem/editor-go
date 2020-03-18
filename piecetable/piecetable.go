package piecetable

import (
	"fmt"
	"io/ioutil"
)

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
	return &table{original: buf, pieces: []piece{{offset: 0, length: len(buf), added: false}}}
}

func NewFromFile(filename string) (*table, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(string(f)), nil
}

func (pt *table) Delete(offset, length int) {
	before := piece{offset: 0, length: offset, added: false}
	after := piece{offset: offset + length, length: len(pt.original) - length - offset, added: false}
	pt.pieces = []piece{before, after}
}

func (pt *table) Get() string {
	sequence := ""
	for _, piece := range pt.pieces {
		if piece.added {
			sequence += pt.add[piece.offset : piece.offset+piece.length]
		} else {
			sequence += pt.original[piece.offset : piece.offset+piece.length]
		}

	}
	return sequence
}

func (pt *table) pieceAt(offset int) (pieceIndex int, pieceOffset int, err error) {
	remaining := offset
	for i, piece := range pt.pieces {
		if remaining < piece.length {
			return i, piece.offset + remaining, nil
		}
		remaining -= piece.offset
	}
	return 0, 0, fmt.Errorf("No piece found at offset %v", offset)
}
