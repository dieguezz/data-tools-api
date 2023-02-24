package cifcontroldigit

import (
	"testing"
)

type cif struct {
	str     string
	control string
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
