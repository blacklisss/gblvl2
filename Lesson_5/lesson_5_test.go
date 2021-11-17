package main

import "testing"

func TestSummer(t *testing.T) {
	tables := []struct {
		res int
	}{
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
		{1000},
	}
	for _, table := range tables {
		var sum = new(int)
		*sum = 0

		sum = Summer(sum)

		if *sum != table.res {
			t.Fatalf("expected 1000, recived %d\n", *sum)
		}
	}
}
