package nifgenerators

import (
	"fmt"
	"math/rand"

	nifcontroldigit "github.com/dieguezz/nif-tools/control-digit"
)

type NIFOptions struct {
	UnderAge          bool `json:"underAge,omitempty" bson:",omitempty"`
	LivingOutside     bool `json:"livingOutside,omitempty" bson:",omitempty"`
	ForeignWithoutNIE bool `json:"foreignWithoutNIE,omitempty" bson:",omitempty"`
}

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GenerateNIF(args NIFOptions) string {

	var prefix string = ""
	var body string

	if !args.UnderAge && !args.ForeignWithoutNIE && !args.LivingOutside {
		num := int32(numberFromRange(10000000, 99999999))
		control := nifcontroldigit.GetControlDigit(num)
		body = fmt.Sprintf("%d%s", num, control)
	} else {
		num := int32(numberFromRange(10000000, 99999999))
		control := nifcontroldigit.GetControlDigit(num)
		body = fmt.Sprintf("%d%s", num, control)
	}

	if args.UnderAge {
		prefix = "K"
	}

	if args.LivingOutside {
		prefix = "L"
	}

	if args.ForeignWithoutNIE {
		prefix = "M"
	}

	return fmt.Sprintf("%s%s", prefix, body)
}
