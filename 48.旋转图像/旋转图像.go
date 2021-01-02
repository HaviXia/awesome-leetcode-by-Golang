package _8_旋转图像

/*
	给定一个 n × n 的二维矩阵表示一个图像。
	将图像顺时针旋转 90 度。
	说明：
		你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
*/

/*
官方题解：
	根据规律：对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。

	由于矩阵中的行列从 0 开始计数，因此对于矩阵中的元素 matrix[row][col] 旋转过后的新位置为 matrix'[col][n-row-1]

	时间复杂度：O(N^2)
	空间复杂度：O(N^2)
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	temp := make([][]int, n) //构造辅助空间N^2的临时矩阵
	for i := range temp {
		temp[i] = make([]int, n)
	}
	for i, row := range matrix {
		for j, v := range row {
			temp[j][n-i-1] = v
		}
	}
	copy(matrix, temp) //引用拷贝
}

/*

 */
