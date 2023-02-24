package nifcontroldigit

import (
	"testing"
)

type nif struct {
	number  int
	control string
}

type cif struct {
	str     string
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

func TestGetCIFControlDigit(t *testing.T) {
	nifList := []cif{{str: "U3417318", control: "7"}, {str: "N0175417", control: "E"}, {str: "J4283343", control: "4"}, {str: "W5287701", control: "F"}}

	for _, x := range nifList {
		result := GetCIFControlDigit(x.str)

		if result != x.control {
			t.Errorf("Failed testing %v, expected %s, received %s", x.str, x.control, result)

		}
	}
}
