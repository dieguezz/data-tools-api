package nifvalidator

import (
	"strconv"
	"testing"
)

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
