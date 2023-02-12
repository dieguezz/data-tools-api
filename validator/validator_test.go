package nifvalidator

import (
	"strconv"
	"testing"
)

func TestIsValidNIE(t *testing.T) {
	validNIEs := []string{"x5555555", "y5555555", "z55555555"}
	inValidNIEs := []string{"5555555", "w5555555", "p55555555", "z999999999"}
	validValue := true

	for _, x := range validNIEs {
		result := IsValidNIE(x)
		if validValue != result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}

	for _, x := range inValidNIEs {
		result := IsValidNIE(x)
		if validValue == result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}

}
func TestIsValidNIF(t *testing.T) {
	validNIFs := []string{"51069845", "92331746", "05885424", "5885424", "L11111111", "l11111111", "51069845R"}
	inValidNIFs := []string{"510r9845", "w5555555", "55555555pp", "r22222222"}
	validValue := true

	for _, x := range validNIFs {
		result := IsValidNIF(x)
		if validValue != result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}

	for _, x := range inValidNIFs {
		result := IsValidNIF(x)
		if validValue == result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))

		}
	}

}
