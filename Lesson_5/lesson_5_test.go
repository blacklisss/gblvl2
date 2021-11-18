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
	var sum = new(int)

	for _, table := range tables {
		*sum = 0

		Summer(sum)

		if *sum != table.res {
			t.Fatalf("expected 1000, recived %d\n", *sum)
		}
	}
}
