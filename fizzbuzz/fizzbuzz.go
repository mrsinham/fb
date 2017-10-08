package fizzbuzz

import (
	"context"
	"strconv"
)

// Naive demonstrate that the naive version can performs really well
// The key is to reduce allocations by allocation the slice with the final capacity
// the backing array dont has to be reallocated many time before being of the final size
func Naive(ctx context.Context, string1 string, string2 string, int1 int, int2 int, limit int) (res []string, err error) {

	// allocation magic here
	res = make([]string, limit)
	for i := 1; i <= limit; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			var s string
			if i%int1 == 0 {
				s = string1
			}
			if i%int2 == 0 {
				s += string2
			}

			if s == "" {
				s = strconv.FormatInt(int64(i), 10)
			}

			// could use append, the price is the same
			res[i-1] = s
		}

	}

	return
}

// Switch demonstrate the ability to use conditionnal switch cases
// but is slower than the Naive version
func Switch(string1 string, string2 string, int1 int, int2 int, limit int) (res []string) {

	res = make([]string, 0, limit)
	concat := string1 + string2

	var s string
	for i := 1; i <= limit; i++ {

		switch {
		case i%(int1*int2) == 0:
			s = concat
		case i%int1 == 0:
			s = string1
		case i%int2 == 0:
			s = string2
		default:
			s = strconv.FormatInt(int64(i), 10)
		}

		res = append(res, s)
	}

	return
}
