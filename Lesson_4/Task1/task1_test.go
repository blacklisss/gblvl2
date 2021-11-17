package main

import (
	"testing"
)

func TestCheckSum(t *testing.T) {

	for i := 0; i < 10; i++ {
		var sum = new(int)
		*sum = 0

		sum = Summer(sum)
		if *sum != 1000 {
			t.Errorf("expected 1000, but recived %d\n", *sum)
		}
	}
}
