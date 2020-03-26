package piecetable

import (
	"fmt"
	"io/ioutil"
	"log"
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

func (pt *table) Delete(offset, length int) error {
	firstIndex, firstOffset, err := pt.pieceAt(offset)
	if err != nil {
		return err
	}
	lastIndex, lastOffset, err := pt.pieceAt(offset + length)
	if err != nil {
		return err
	}
	if firstIndex == lastIndex {
		piece := &pt.pieces[firstIndex]
		if firstOffset == piece.offset {
			piece.offset += length
			piece.length -= length
			return err
		}
		if lastOffset == piece.offset+piece.length {
			piece.length -= length
			return err
		}
	}

	first := pt.pieces[firstIndex]
	last := pt.pieces[lastIndex]
	deleted := []piece{
		{added: first.added, offset: first.offset, length: firstOffset - first.offset},
		{added: last.added, offset: lastOffset, length: last.length - (lastOffset - last.offset)},
	}
	filtered := []piece{}
	for _, piece := range deleted {
		if piece.length > 0 {
			filtered = append(filtered, piece)
		}
	}
	newPieces := append(pt.pieces[:firstIndex], filtered...)
	newPieces = append(newPieces, pt.pieces[lastIndex-firstIndex+1:]...)
	pt.pieces = newPieces
	return err
}

func (pt *table) Insert(new string, offset int) error {
	addOffset := len(pt.add)
	pt.add += new
	pieceIndex, pieceOffset, err := pt.pieceAt(offset)
	log.Printf("piece at: %v / %v", pieceIndex, pieceOffset)
	if err != nil {
		return err
	}
	original := &pt.pieces[pieceIndex]
	if endOfAddBuffer(original, pieceOffset, addOffset) {
		original.length += len(new)
		return nil
	}

	inserted := []piece{
		{added: original.added, offset: original.offset, length: pieceOffset - original.offset},
		{added: true, offset: addOffset, length: len(new)},
		{added: original.added, offset: pieceOffset, length: original.length - (pieceOffset - original.offset)},
	}
	filtered := []piece{}
	for _, piece := range inserted {
		if piece.length > 0 {
			filtered = append(filtered, piece)
		}
	}
	newPieces := append(pt.pieces[:pieceIndex], filtered...)
	newPieces = append(newPieces, pt.pieces[pieceIndex+1:]...)
	pt.pieces = newPieces
	log.Printf("sequence: %v", pt.Get())
	log.Printf("add buffer: %s", pt.add)
	log.Printf("pieces: %v", pt.pieces)

	return err
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
		if remaining <= piece.length {
			return i, piece.offset + remaining, nil
		}
		remaining -= piece.offset
	}
	return 0, 0, fmt.Errorf("No piece found at offset %v", offset)
}

func endOfAddBuffer(original *piece, pieceOffset, addOffset int) bool {
	return original.added && pieceOffset == (original.offset+original.length) && (original.offset+original.length == addOffset)
}
