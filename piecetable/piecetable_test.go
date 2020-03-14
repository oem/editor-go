package piecetable

import "testing"

func TestNew(t *testing.T) {
	expected := "moo"
	got := New(expected)
	if got.original != expected {
		t.Errorf("Expected new piecetable to have the original buffer set to moo, got %v", got.original)
	}

	if len(got.pieces) != 1 {
		t.Errorf("Expected one piece, got %v pieces", len(got.pieces))
	}
}

func TestNewFirstPiece(t *testing.T) {
	expected := "moo"
	got := New(expected)
	piece := got.pieces[0]

	if piece.length != len(expected) {
		t.Errorf("Expected first piece to have length of original, got %v instead", piece.length)
	}

	if piece.added {
		t.Errorf("Expected first piece to point to original buffer, pointing to add buffer instead")
	}

	if piece.offset != 0 {
		t.Errorf("Expected first piece to have an offset of 0, got %v instead", piece.offset)
	}
}

func TestNewFromFile(t *testing.T) {
	expected := "moo in file\n"
	got, err := NewFromFile("test.txt")
	if err != nil {
		t.Errorf("expected success, got error %v instead", err)
	}

	if got.original != expected {
		t.Errorf("Expected %v, got %v", expected, got.original)
	}

	if len(got.pieces) != 1 {
		t.Errorf("Expected one piece, got %v pieces", len(got.pieces))
	}
}
