package arrays

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArraysHandler struct {
	Service *ArrayService
}

func (h *ArraysHandler) ContainsDuplicates(c *gin.Context) {
	var numbers []int
	if err := c.ShouldBindJSON(&numbers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Service.ContainsDuplicates(numbers)

	c.JSON(http.StatusOK, result)
}

type SudokuBoard struct {
	Board [][]string `json:board`
}

func (h *ArraysHandler) IsValidSudoku(c *gin.Context) {
	var sudokuBoard SudokuBoard
	if err := c.ShouldBindJSON(&sudokuBoard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board := make([][]byte, len(sudokuBoard.Board))
	for i, row := range sudokuBoard.Board {
		board[i] = make([]byte, len(row))
		for j, cell := range row {
			board[i][j] = cell[0]
		}
	}

	result := h.Service.IsValidSudoku(board)
	fmt.Println(result)
	c.JSON(http.StatusOK, result)
}
