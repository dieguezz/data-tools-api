package nifcontroldigit

import (
	"testing"
)

type nif struct {
	number  int
	control string
}

func TestGetControlDigit(t *testing.T) {
	nifList := []nif{{number: 41509090, control: "R"}, {number: 40758223, control: "S"}, {number: 37909898, control: "X"}, {number: 74255828, control: "Y"}}

	for _, x := range nifList {
		result := GetControlDigit(int32(x.number))

		if result != x.control {
			t.Fatal("Failed testing")
		}
	}
}
