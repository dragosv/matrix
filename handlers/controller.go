package handlers

import (
	"github.com/dragosv/matrix/commands"
	"github.com/dragosv/matrix/processing"
)

type MatrixController struct {
	operations commands.MatrixOperations
	reader     processing.MatrixReader
}
