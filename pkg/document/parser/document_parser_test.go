package documentparser

import (
	"reflect"
	"testing"
)

type nifCandidate struct {
	nif      string
	expected ParsedDcocument
}

func TestGetParsedDcocument(t *testing.T) {
	items := []nifCandidate{{nif: "74255828Y", expected: ParsedDcocument{Number: 74255828, Kind: "NIF", Control: "Y"}}, {nif: "74255828Y", expected: ParsedDcocument{Number: 74255828, Kind: "NIF", Control: "Y"}}}

	for _, x := range items {
		result, _ := GetParsedDocument(x.nif)
		if !reflect.DeepEqual(x.expected, result) {
			t.Fatal("Failed testing")
		}
	}
}
