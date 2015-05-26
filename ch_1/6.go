package ctci_chapter1

type coordPair struct {
	row int
	col int
}

type itrDef struct {
	start int
	lim int
	dir int
}

func RotateImage(input [][]int32) [][]int32 {
	//new plan: when swapping, set up threads
	//each thread comunicates via a channel
	//to give next x,y coordinate
	//main thread receives these values, then calls swap on them

	quadrantDefs := []coordPair {
		{0,0},
		{0,1},
		{1,1},
		{1,0},
	}

	inLen := len(input)

	//now, swap in place:
	//q[0],q[1]
	//q[0],q[2]
	//q[0],q[3]
	for q:=1; q < 4; q++ {
		c1, c2 := make(chan coordPair), make(chan coordPair)
		go quadrantItr(quadrantDefs[0].row, quadrantDefs[0].col, inLen, c1)
		go quadrantItr(quadrantDefs[q].row, quadrantDefs[q].col, inLen, c2)

		iterations := int(inLen/2)
		iterations *= (iterations+inLen%2)
		for i:=0; i < iterations; i++ {
			old := <-c1
			knew := <-c2
			swap(old.row, old.col, knew.row, knew.col, input[:][:])
		}
	}

	return input
}

func quadrantItr(row int, col int, length int, output chan<- coordPair) {
	colMajor := (row-col == 0)
	defs := make([]itrDef,2)
	defs[0] = makeDef(row, length, !colMajor)
	defs[1] = makeDef(col, length, colMajor)

	if colMajor {
		for j := defs[1].start; j != defs[1].lim; j+=defs[1].dir {
			for i := defs[0].start; i != defs[0].lim; i+=defs[0].dir {
				output <- coordPair{i, j}
			}
		}
	} else {
		for i := defs[0].start; i != defs[0].lim; i+=defs[0].dir {
			for j := defs[1].start; j != defs[1].lim; j+=defs[1].dir {
				output <- coordPair{i, j}
			}
		}
	}
}

func makeDef(in int, length int, isMajor bool) itrDef {
	var ret itrDef
	if in == 0 {
		ret.start = 0
		ret.lim = int(length/2)
		if isMajor {
			ret.lim += length%2
		}
		ret.dir = 1
	} else {
		ret.start = length-1
		ret.lim = int(length/2)-1
		if !isMajor {
			ret.lim += length%2
		}
		ret.dir = -1
	}
	return ret
}

func swap(x1 int, y1 int, x2 int, y2 int, array [][]int32) {
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
	array[x2][y2] = array[x1][y1] ^ array[x2][y2]
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
}
