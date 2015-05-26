package ctci_chapter1

import (
	"testing"
	"fmt"
)

func TestRotate(t *testing.T) {
	cases := []struct {
		in [][]int32
		want [][]int32
	}{
		{
			[][]int32 {
				{0,1,2,3},
				{4,5,6,7},
				{8,9,10,11},
				{12,13,14,15},
			},
			[][]int32 {
				{12,8,4,0},
				{13,9,5,1},
				{14,10,6,2},
				{15,11,7,3},
			},
		},
		{
			[][]int32 {
				{1,2,3},
				{4,5,6},
				{7,8,9},
			},
			[][]int32 {
				{7,4,1},
				{8,5,2},
				{9,6,3},
			},
		},
	}

	for _, c := range cases {
		got := RotateImage(c.in)
		if sliceEqual32(got[:][:], c.want[:][:]) == false {
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

func sliceEqual32(a [][]int32, b [][]int32) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
