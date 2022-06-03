package main

type NumMatrix struct {
	dp [][]int
}

func Constructor(matrix [][]int) NumMatrix {

	var data NumMatrix

	data.dp = make([][]int, len(matrix)+1)
	for i := range data.dp {
		data.dp[i] = make([]int, len(matrix[0])+1)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			data.dp[i+1][j+1] = data.dp[i+1][j] + data.dp[i][j+1] + matrix[i][j] - data.dp[i][j]
		}
	}

	return data
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {

	return this.dp[row2+1][col2+1] - this.dp[row2+1][col1] - this.dp[row1][col2+1] +
		this.dp[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
