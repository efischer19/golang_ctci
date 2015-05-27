package ctci_chapter1

import (
	"testing"
	"fmt"
)

func TestZeroRowCol(t *testing.T) {
	cases := []struct {
		in [][]int
		want [][]int
	}{
		{
			[][]int {
				{0,1,2,3},
				{4,5,6,7},
				{8,9,0,11},
				{12,13,14,15},
			},
			[][]int {
				{0,0,0,0},
				{0,5,0,7},
				{0,0,0,0},
				{0,13,0,15},
			},
		},
		{
			[][]int {
				{1,2,3},
				{4,0,6},
				{7,8,9},
			},
			[][]int {
				{1,0,3},
				{0,0,0},
				{7,0,9},
			},
		},
	}

	for _, c := range cases {
		got := ZeroRowCol(c.in)
		if sliceEqual(got[:][:], c.want[:][:]) == false {
			t.Errorf("Error! Results follow")
			for i := 0; i < len(got); i++ {
				fmt.Println(got[i])
			}
			for i := 0; i < len(c.want); i++ {
				fmt.Println(c.want[i])
			}
		}
	}
}

func sliceEqual(a [][]int, b [][]int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
