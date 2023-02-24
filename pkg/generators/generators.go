package nifgenerators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	nifcontroldigit "github.com/dieguezz/nif-tools/pkg/control-digit"
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
		control := nifcontroldigit.GetNIFControlDigit(num)
		body = fmt.Sprintf("%d%s", num, control)
	} else {
		num := int32(numberFromRange(1000000, 9999999))
		control := nifcontroldigit.GetNIFControlDigit(num)
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

func GenerateNIE() string {
	num := numberFromRange(1000000, 9999999)
	body := fmt.Sprintf("%d", num)

	in := []string{"X", "Y", "Z"}
	prefix := in[rand.Intn(len(in))]
	prefixMapping := map[string]string{"X": "", "Y": "1", "Z": "2"}
	base := fmt.Sprintf("%s%d", prefixMapping[prefix], num)
	baseNum, err := strconv.Atoi(base)

	if err != nil {
		log.Fatalf("Failed generating NIE %v", err)
	}

	control := nifcontroldigit.GetNIFControlDigit(int32(baseNum))

	return fmt.Sprintf("%s%s%s", prefix, body, control)
}

func GenerateCIF() string {
	availableLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "Q", "S", "K", "L", "M", "X"}
	num := numberFromRange(1000000, 9999999)
	letter := availableLetters[rand.Intn(len(availableLetters))]
	control := nifcontroldigit.GetCIFControlDigit(fmt.Sprintf("%s%v", letter, num))
	return fmt.Sprintf("%v%v%v", letter, num, control)
}
