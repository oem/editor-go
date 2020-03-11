package piecetable

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
