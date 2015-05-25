package ctci_chapter1

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
		{1,1},
		{-1,1},
		{-1,-1},
		{1,-1},
	}

	//now, swap in place:
	//q[0],q[1]
	//q[0],q[2]
	//q[0],q[3]
}

func swap(x1 int, y1 int, x2 int, y2 int, array [][]int32) {
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
	array[x2][y2] = array[x1][y1] ^ array[x2][y2]
	array[x1][y1] = array[x1][y1] ^ array[x2][y2]
}
