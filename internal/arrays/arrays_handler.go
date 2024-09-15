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

func (h *ArraysHandler) ValidAnagram(c *gin.Context) {
	var strings struct {
		S string `json:"s"`
		T string `json:"t"`
	}
	if err := c.ShouldBindJSON(&strings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.Service.ValidAnagram(strings.S, strings.T)
	c.JSON(http.StatusOK, result)
}
