package float

import (
	"testing"
)

// Tes ...
type Tes struct {
	ID     int
	Name   string
	Amount float64
}

// Tes2 ...
type Tes2 struct {
	Angka float64
}

const (
	precision = 5
)

func TestFilterFloat(t *testing.T) {
	var tes []Tes
	tes = append(tes, Tes{
		ID:     1,
		Name:   "Jannes",
		Amount: 10290.1231231231223,
	})
	tes = append(tes, Tes{
		ID:     2,
		Name:   "Santoso",
		Amount: 303.341241245,
	})
	FilterFloat(&tes, precision)

	if tes[0].Amount != 10290.12312 {
		t.Fatal(10290.12312, "Should be", tes[0].Amount)
	}

	if tes[1].Amount != 303.34124 {
		t.Fatal(303.34124, "Should be", tes[1].Amount)
	}

	tes2 := &Tes2{
		Angka: 20.2038209384,
	}
	FilterFloat(&tes2, precision)

	if (*tes2).Angka != 20.20382 {
		t.Fatal(20.20382, "Should be", (*tes2).Angka)
	}
}
