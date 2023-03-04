package cifgenerators

import (
	"fmt"
	"math/rand"
	"sync"

	cifcontroldigit "github.com/dieguezz/nif-tools/pkg/cif/control-digit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func numberFromRange(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func GenerateCIF() string {
	availableLetters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "P", "Q", "S", "K", "L", "M", "X"}
	num := numberFromRange(1000000, 9999999)
	letter := availableLetters[rand.Intn(len(availableLetters))]
	control := cifcontroldigit.GetCIFControlDigit(fmt.Sprintf("%s%v", letter, num))
	return fmt.Sprintf("%v%v%v", letter, num, control)
}


func GenerateCIFs(amount int) ([]string, error) {
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
			nif := GenerateCIF()
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