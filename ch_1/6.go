package ctci_chapter1

import "fmt"

type coordPair struct {
	x int
	y int
}

func RotateImage(input [][]int32) [][]int32 {
	//new plan: when swapping, set up threads
	//each thread comunicates via a channel
	//to give next x,y coordinate
	//main thread receives these values, then calls swap on them

	//for these defs:
	//in form {x,y}
	//1 means: start at 0, increment to limit
	//-1 means: start at max, decrement to limit
	//if same sign, x is major axis and has limit 1 further
	//opposite sign: y is major axis and limit extends 1 further
	quadrantDefs := []struct {
		x int
		y int
	} {
		{1,1},
		{-1,1},
		{-1,-1},
		{1,-1},
	}

	var quadLen = len(input)
	//now, swap in place:
	//q[0],q[1]
	//q[0],q[2]
	//q[0],q[3]
	for q:=1; q < 4; q++ {
		c1, c2 := make(chan coordPair), make(chan coordPair)
		go quadrantItr(quadrantDefs[0].x, quadrantDefs[0].y, quadLen, c1)
		go quadrantItr(quadrantDefs[q].x, quadrantDefs[q].y, quadLen, c2)

		for i:=0; i < int(quadLen/2)*int(quadLen/2 + quadLen%2); i++ {
			old := <-c1
			knew := <-c2
			fmt.Printf("Swapping %d,%d with %d,%d\n",old.x,old.y,knew.x,knew.y);
			swap(old.x, old.y, knew.x, knew.y, input[:][:])
		}
	}

	return input
}

func quadrantItr(x int, y int, max int, output chan<- coordPair) {
	var out, in int
	if x*y > 0 {
		//x is outer
		if x > 0 {
			out = 0
		} else {
			out = max-1
		}

		if y > 0 {
			in = 0
		} else {
			in = max-1
		}
	} else {
		//y is outer
		if y > 0 {
			out = 0
		} else {
			out = max-1
		}

		if x > 0 {
			in = 0
		} else {
			in = max-1
		}
	}

	lim := int(max/2)
	limMod := max%2

	outDir := out/(max-1)	//outDir is 1 (if desc) or 0 (if asc)
	outLim := lim-outDir
	outDir += (outDir-1)	//outDir is now either 1 or -1
	outDir *= -1			//get the sign right

	inDir := in/(max-1)
	inLim := lim-inDir
	inDir += (inDir-1)		//inDir is now either 1 or -1
	inDir *= -1


	for ; out != outLim+(limMod*outDir); out += outDir {
		for inTemp := in; inTemp != inLim; inTemp += inDir {
			if x*y > 0 {
				output <- coordPair{inTemp,out}
			} else {
				output <- coordPair{out,inTemp}
			}
		}
	}
}

func swap(x1 int, y1 int, x2 int, y2 int, array [][]int32) {
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
	array[x2][y2] = array[x1][y1] ^ array[x2][y2]
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
}
