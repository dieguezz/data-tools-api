package documentvalidator

import (
	"context"

	cifvalidator "github.com/dieguezz/nif-tools/pkg/cif/validators"
	nievalidator "github.com/dieguezz/nif-tools/pkg/nie/validators"
	nifvalidator "github.com/dieguezz/nif-tools/pkg/nif/validators"
)

func IsValidCIFNIENIF(str string) bool {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan bool)

	go func() {
		ch <- cifvalidator.IsValidCIF(str)
	}()

	go func() {
		ch <- nifvalidator.IsValidNIF(str)
	}()

	go func() {
		ch <- nievalidator.IsValidNIE(str)
	}()

	for i := 0; i < 3; i++ {
		select {
		case res := <-ch:
			if res {
				cancel()
				return true
			}
		case <-ctx.Done():
			return false
		}
	}

	return false
}
