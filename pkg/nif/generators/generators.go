package nifgenerators

import (
	"fmt"
	"math/rand"
	"sync"

	nifcontroldigit "github.com/dieguezz/nif-tools/pkg/nif/control-digit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type NIFOptions struct {
	UnderAge          bool  `json:"underAge,omitempty" bson:",omitempty"`
	LivingOutside     bool  `json:"livingOutside,omitempty" bson:",omitempty"`
	ForeignWithoutNIE bool  `json:"foreignWithoutNIE,omitempty" bson:",omitempty"`
	Amount            int32 `json:"amount,omitempty" bson:",omitempty"`
}

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GeneratedNIFs(args NIFOptions) ([]string, error) {
	var wg sync.WaitGroup
	results := make(chan string, args.Amount)
	limit := 100000
	if int(args.Amount) > limit {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Max limit reached: 'amount' should not be bigger than %v", limit))
	}
	for i := 0; i < int(args.Amount); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nif := GenerateNIF(args)
			results <- nif
		}()
	}

	wg.Wait()
	close(results)

	var res []string

	for nif := range results {
		res = append(res, nif)
	}
	return res, nil
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
