package nifparser

import (
	"reflect"
	"testing"
)

type nifCandidate struct {
	nif      string
	expected ParsedNIF
}

func TestGetParsedNIF(t *testing.T) {
	items := []nifCandidate{{nif: "74255828Y", expected: ParsedNIF{Number: 74255828, Kind: "NIF", Control: "Y"}}, {nif: "74255828Y", expected: ParsedNIF{Number: 74255828, Kind: "NIF", Control: "Y"}}}

	for _, x := range items {
		result, _ := GetParsedNIF(x.nif)
		if !reflect.DeepEqual(x.expected, result) {
			t.Fatal("Failed testing")
		}
	}
}
