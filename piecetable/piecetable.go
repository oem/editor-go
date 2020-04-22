package piecetable

import (
	"fmt"
	"io/ioutil"
)

// Table exposes a simple piece-table
type Table struct {
	original string
	add      string
	pieces   []piece
}

type piece struct {
	offset int
	length int
	added  bool
}

// New is creating a new piecetable, using the provided string as original buffer.
// Arguments: buf string
// Returns: *Table
func New(buf string) *Table {
	return &Table{original: buf, pieces: []piece{{offset: 0, length: len(buf), added: false}}}
}

// NewFromFile will try to open the file and instantiate a piecetable.
// Arguments: filename string
// Returns: *table, error
func NewFromFile(filename string) (*Table, error) {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return New(string(f)), nil
}

// Delete will remove and rechunk parts of pieces at the offset/length
func (pt *Table) Delete(offset, length int) error {
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
			return nil
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

// Insert adds new content to the add buffer and created pieces that point to the new text
func (pt *Table) Insert(new string, offset int) error {
	addOffset := len(pt.add)
	pt.add += new
	pieceIndex, pieceOffset, err := pt.pieceAt(offset)
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
	newPieces := make([]piece, len(pt.pieces))
	copy(newPieces, pt.pieces[:pieceIndex])
	newPieces = append(newPieces, filtered...)
	pt.pieces = append(newPieces, pt.pieces[pieceIndex+1:]...)

	return err
}

// Get returns the current state of the piece-table
func (pt *Table) Get() string {
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

func (pt *Table) pieceAt(offset int) (pieceIndex int, pieceOffset int, err error) {
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
