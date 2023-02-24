package niegenerators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	niecontroldigit "github.com/dieguezz/nif-tools/pkg/nie/control-digit"
)

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
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

	control := niecontroldigit.CalculateNIELetter(int32(baseNum))

	return fmt.Sprintf("%s%s%s", prefix, body, control)
}
