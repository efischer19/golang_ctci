package ctci_chapter1

func ZeroRowCol(input [][]int) [][]int {
	//first thought: 2-pass, use some extra memory
	n := len(input)
	m := len(input[0])

	var rowZeros = make([]bool, n)
	var colZeros = make([]bool, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if input[i][j] == 0 {
				rowZeros[i] = true
				colZeros[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if rowZeros[i] || colZeros[j] {
				input[i][j] = 0
			}
		}
	}

	return input
}
