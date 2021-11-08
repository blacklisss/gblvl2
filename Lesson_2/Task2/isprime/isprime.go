// Copyright 2021 Romanov Yuri. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package isprime for determining prime numbers
// func IsPrime(n int) (ret bool)

package isprime

import "math"

// IsPrime get int number and returns boolean
func IsPrime(n int) (ret bool) {
	x := int(math.Sqrt(float64(n)))
	for i := 2; i <= x; i++ {
		if n%i == 0 {
			return ret
		}
	}

	return !ret
}
