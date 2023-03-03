package niegenerators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"

	niecontroldigit "github.com/dieguezz/nif-tools/pkg/nie/control-digit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GeneratedNIEs(amount int) ([]string, error) {
	var wg sync.WaitGroup
	results := make(chan string, amount)
	limit := 100000
	if int(amount) > limit {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("Max limit reached: 'amount' should not be bigger than %v", limit))
	}
	for i := 0; i < int(amount); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			nie := GenerateNIE()
			results <- nie
		}()
	}

	wg.Wait()
	close(results)

	var res []string

	for nie := range results {
		res = append(res, nie)
	}
	return res, nil
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
