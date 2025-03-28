//# https://theweeklychallenge.org/blog/perl-weekly-challenge-300/
/*#

Task 2: Nested Array

Submitted by: [55]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an array of integers, @ints of length n containing
   permutation of the numbers in the range [0, n - 1].

   Write a script to build a set, set[i] = ints[i], ints[ints[i]],
   ints[ints[ints[i]]], ..., subjected to the following rules:
1. The first element in set[i] starts with the selection of elements ints[i].
2. The next element in set[i] should be ints[ints[i]], and then ints[ints[ints[i
]]], and so on.
3. We stop adding right before a duplicate element occurs in set[i].

   Return the longest length of a set set[i].

Example 1

Input: @ints = (5, 4, 0, 3, 1, 6, 2)
Output: 4

ints[0] = 5
ints[1] = 4
ints[2] = 0
ints[3] = 3
ints[4] = 1
ints[5] = 6
ints[6] = 2

One of the longest sets set[k]:
set[0] = {ints[0], ints[5], ints[6], ints[2]} = {5, 6, 2, 0}

Example 2

Input: @ints = (0, 1, 2)
Output: 1
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 22^nd December
   2024.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
#*/
//# solution by pokgopun@gmail.com

package main

import (
	"io"
	"os"

	"github.com/google/go-cmp/cmp"
)

type ints []int

func (is ints) na() int {
	seen := make(map[int]bool)
	mx := 0
	for i := range len(is) {
		n := is[i]
		for !seen[n] {
			seen[n] = true
			n = is[n]
		}
		l := len(seen)
		if mx < l {
			mx = l
		}
		clear(seen)
	}
	return mx
}

func main() {
	for _, data := range []struct {
		input  ints
		output int
	}{
		{ints{5, 4, 0, 3, 1, 6, 2}, 4},
		{ints{0, 1, 2}, 1},
	} {
		io.WriteString(os.Stdout, cmp.Diff(data.input.na(), data.output)) // blank if ok, otherwise show the difference
	}
}
