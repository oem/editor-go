package piecetable

import "testing"

func TestNew(t *testing.T) {
	expected := "moo"
	got := New(expected)
	if got.original != expected {
		t.Errorf("Expected new piecetable to have the original buffer set to moo, got %v", got.original)
	}
}

func TestNewFromFile(t *testing.T) {
	expected := "moo in file\n"
	got, err := NewFromFile("test.txt")
	if err != nil {
		t.Errorf("expected success, got error %v instead", err)
		return
	}
	if got.original != expected {
		t.Errorf("Expected %v, got %v", expected, got.original)
	}
}
