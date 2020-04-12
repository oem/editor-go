package piecetable

import (
	"testing"
)

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

func TestGet(t *testing.T) {
	want := "moo foo goo"
	pt := New(want)
	got := pt.Get()

	if got != want {
		t.Errorf("Expected '%s', got '%s' instead", want, got)
	}
}

func TestDeletingOnce(t *testing.T) {
	pt := New("moo foo goo")
	pt.Delete(4, 4)
	if len(pt.pieces) != 2 {
		t.Errorf("Expected two pieces after delete, got %v", len(pt.pieces))
	}

	if pt.pieces[0].length != 4 {
		t.Errorf("Expected first piece length to be until start of deletion(5), but was %v", pt.pieces[0].length)
	}
	if pt.pieces[1].length != 3 {
		t.Errorf("Expected second piece length to be until end of original buffer(3), but was %v", pt.pieces[1].length)
	}

	if pt.pieces[0].offset != 0 {
		t.Errorf("Expected first piece to have an offset of zero, got %v", pt.pieces[0].offset)
	}
	if pt.pieces[1].offset != 8 {
		t.Errorf("Expected first piece to have an offset of 8, got %v", pt.pieces[1].offset)
	}

	if pt.Get() != "moo goo" {
		t.Errorf("Expected 'moo goo' after Delete, got '%s' instead", pt.Get())
	}
}

func TestDeletingTwice(t *testing.T) {
	pt := New("moo foo goo")
	err := pt.Delete(4, 4)
	if err != nil {
		t.Errorf("Expected Delete(4, 4) to work but got error: %v", err)
	}
	err = pt.Delete(0, 1)
	if err != nil {
		t.Errorf("Expected Delete(0, 1) to work but got error: %v", err)
	}
	want := "oo goo"
	got := pt.Get()

	if got != want {
		t.Errorf("Expected '%s' after two deletes, got '%s'", want, got)
	}
}

func TestShorteningLastPiece(t *testing.T) {
	pt := New("foo")
	err := pt.Delete(2, 1)
	if err != nil {
		t.Errorf("Expected deleting to be successful but got error: %v", err)
	}
	want := "fo"
	got := pt.Get()
	if got != want {
		t.Errorf("Expected %s, got %s instead", want, got)
	}
}

func TestInsertingOnce(t *testing.T) {
	pt := New("0123456789")
	err := pt.Insert("MOO", 3)
	if err != nil {
		t.Errorf("Expected Insert('MOO ', 3) to be successful, got %v instead", err)
	}
	want := "012MOO3456789"
	got := pt.Get()

	if got != want {
		t.Errorf("Expected %s, got %s instead", want, got)
	}
}

func TestInsertingMultipleTimes(t *testing.T) {
	pt := New("0123456789")
	err := pt.Insert("MOO", 9)
	if err != nil {
		t.Errorf("Expected Insert('MOO ', 9) to be successful, got %v instead", err)
	}
	want := "012345678MOO9"
	got := pt.Get()
	if got != want {
		t.Errorf("Expected %s, got %s instead", want, got)
	}

	err = pt.Insert("FOO", 3)
	if err != nil {
		t.Errorf("Expected Insert('FOO ', 3) to be successful, got %v instead", err)
	}
	want = "012FOO345678MOO9"
	got = pt.Get()

	if got != want {
		t.Errorf("Expected %s, got %s instead", want, got)
	}
}

func TestInsertingEmpty(t *testing.T) {
	want := "moo goo"
	pt := New(want)
	err := pt.Insert("", 1)
	if err != nil {
		t.Errorf("Expected insert to be successful, got %v instead", err)
	}
	got := pt.Get()
	if got != want {
		t.Errorf("Expected %s, got %s instead", want, got)
	}
}

func TestInsertOutOfBounds(t *testing.T) {
	pt := New("moo")
	err := pt.Insert("foo", 100)
	if err == nil {
		t.Errorf("Expected error when trying to insert out of bounds, succeeded instead")
	}
}

func TestInsertToAddBufferPiece(t *testing.T) {
	pt := New("")
	err := pt.Insert("foo", 0)
	if err != nil {
		t.Errorf("Expected inserting to be successful, got error %v instead", err)
	}
	err = pt.Insert("moo", 3)
	if err != nil {
		t.Errorf("Expected inserting to be successful, got error %v instead", err)
	}
	want := "foomoo"
	got := pt.Get()
	if got != want {
		t.Errorf("Expected %v, got %s instead", want, got)
	}
}

func TestInsertsAndDeletes(t *testing.T) {
	pt := New("")
	err := pt.Insert("foo", 0)
	if err != nil {
		t.Errorf("Expected inserting to be successful, got error %v instead", err)
	}
	err = pt.Delete(2, 1)
	if err != nil {
		t.Errorf("Expected deleting to be successful, got error %v instead", err)
	}
	err = pt.Insert("xx", 1)
	if err != nil {
		t.Errorf("Expected inserting to be successful, got error %v instead", err)
	}
	got := pt.Get()
	want := "fxxo"
	if got != want {
		t.Errorf("Expected %s, got %s", want, got)
	}
}
