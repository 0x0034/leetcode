package twoDimensionalPrefixSum

type NumMatrix struct {
	s [][]int
}

func Constructor(matrxi [][]int) NumMatrix {
	m,n := len(matrxi),len(matrxi[0])
	s := make([][]int,m+1)
	for i := range s {
		s[i] = make([]int,n+1)
	}

	for i , row := range matrxi {
		for j , v := range row {
			s[i+1][j+1] = s[i][j+1] + s[i+1][j] - s[i][j] + v
		}
	}
	return NumMatrix{s}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.s[row2+1][col2+1] - this.s[row2+1][col1] - this.s[row1][col2+1] + this.s[row1][col1]
}
