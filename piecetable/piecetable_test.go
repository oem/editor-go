package piecetable

import "testing"

func TestNew(t *testing.T) {
	new := New("moo")
	if new.original != "moo" {
		t.Errorf("Expected new piecetable to have the original buffer set to moo, got %v", new.original)
	}
}
