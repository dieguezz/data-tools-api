package nifvalidator

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
func TestIsValidNIF(t *testing.T) {
	validNIFs := []string{"51069845R", "92331746V", "05885424T", "05885424T", "5555555c"}
	inValidNIFs := []string{"510r9845", "w5555555", "55555555pp", "r22222222", "5555555e"}
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

func TestLooksLikeNIF(t *testing.T) {
	validNIFs := []string{"51069845", "92331746", "05885424", "5885424", "L11111111", "l11111111", "51069845R"}
	inValidNIFs := []string{"510r9845", "w5555555", "55555555pp", "r22222222"}
	validValue := true

	for _, x := range validNIFs {
		result := LooksLikeNIF(x)
		if validValue != result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))
		}
	}

	for _, x := range inValidNIFs {
		result := LooksLikeNIF(x)
		if validValue == result {
			t.Fatalf("Failed testing %s, received %s", x, strconv.FormatBool(result))

		}
	}
}
