package commands

func Invert(records [][]int) (string, error) {
	var invertedRecords [][]int

	invertedRecords = make([][]int, len(records))

	for i := range records {
		invertedRecords[i] = make([]int, len(records))
	}

	for i, record := range records {
		for j, value := range record {
			invertedRecords[j][i] = value
		}
	}

	return Echo(invertedRecords)
}
