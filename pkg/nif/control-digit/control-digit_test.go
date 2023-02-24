package nifcontroldigit

import (
	"testing"
)

type nif struct {
	number  int
	control string
}

func TestGetNIFControlDigit(t *testing.T) {
	nifList := []nif{{number: 41509090, control: "R"}, {number: 40758223, control: "S"}, {number: 37909898, control: "X"}, {number: 74255828, control: "Y"}}

	for _, x := range nifList {
		result := GetNIFControlDigit(int32(x.number))

		if result != x.control {
			t.Errorf("Failed testing %v, expected %s, received %s", x.number, x.control, result)
		}
	}
}
