package commands

type Matrix struct {
}

type MatrixOperations interface {
	Echo(matrix [][]int) (string, error)
	Flatten(matrix [][]int) (string, error)
	Invert(matrix [][]int) (string, error)
	Multiply(matrix [][]int) (string, error)
	Sum(matrix [][]int) (string, error)
	New(matrix [][]int) (string, error)
}
