package commands

//Invert returns the matrix as a string in matrix format where the columns and rows are inverted
func (m *Matrix) Invert(matrix [][]int) (string, error) {
	var invertedRecords [][]int

	invertedRecords = make([][]int, len(matrix))

	for i := range matrix {
		invertedRecords[i] = make([]int, len(matrix))
	}

	for i, row := range matrix {
		for j, value := range row {
			invertedRecords[j][i] = value
		}
	}

	return m.Echo(invertedRecords)
}
