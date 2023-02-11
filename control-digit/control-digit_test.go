package controldigit

import (
	"strconv"
	"testing"
	"reflect"
)

// TestGetIsNIE calls controldigit.GetIsNIE with a string, checking for a valid return value
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

type nif struct {
	number int
	control string
}

func TestGetControlDigit(t *testing.T) {
	nifList := []nif{nif{number: 41509090, control: "R" }, nif{number: 40758223, control: "S"}, nif{number: 37909898, control: "X"}, nif{number: 74255828, control: "Y"}}

	for _, x := range nifList {
		result := GetControlDigit(int32(x.number))

		if result != x.control {
			t.Fatal("Failed testing" )
		}
	}
}

type nifCandidate struct {
	nif string
	expected ParsedNIF
}

func TestGetParsedNIF(t *testing.T) {
	items := []nifCandidate{nifCandidate{nif: "74255828Y", expected: ParsedNIF{ Number: 74255828, Kind: "NIF", Control: "Y"}},nifCandidate{nif: "74255828Y", expected: ParsedNIF{ Number: 74255828, Kind: "NIF", Control: "Y"}}}
		
	for _, x := range items {
		result := GetParsedNIF(x.nif)
		if !reflect.DeepEqual(x.expected, result) {
			t.Fatal("Failed testing" )
		}
	}
}
