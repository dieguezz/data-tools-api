package nievalidator

import (
	"strconv"
	"testing"
)

func TestIsValidNIE(t *testing.T) {
	validNIEs := []string{"Z2321476E", "Y0204374X", "Z8000342G", "X4803950D"}
	inValidNIEs := []string{"5555555", "w5555555", "p55555555", "z5040059M"}
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

func TestLooksLikeNIE(t *testing.T) {
	validNIEs := []string{"Z2321476E", "Y0204374X", "Z8000342G", "z5040059M"}
	inValidNIEs := []string{"5555555", "w5555555", "p55555555", "z999999999"}
	validValue := true

	for _, x := range validNIEs {
		result := LooksLikeNIE(x)
		if validValue != result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}

	for _, x := range inValidNIEs {
		result := LooksLikeNIE(x)
		if validValue == result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}
}
