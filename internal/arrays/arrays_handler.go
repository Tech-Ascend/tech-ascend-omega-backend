package arrays

import (
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

func (h *ArraysHandler) IsValidSudoku(c *gin.Context) {
	var board [][]byte
	if err := c.ShouldBindBodyWithJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	result := h.Service.IsValidSudoku(board)
	c.JSON(http.StatusOK, result)
}
