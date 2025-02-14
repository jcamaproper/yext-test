package handler

import (
	customErrors "app/error"
	"app/model"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SortHandler(c *gin.Context) {
	var req model.SortRequestPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidJSON.Error()})
		return
	}

	if len(req.SortKeys) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidSortKeys.Error()})
		return
	}

	if len(req.Payload) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": customErrors.ErrInvalidPayload.Error()})
		return
	}

	sortedPayload, err := service.SortPayload(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, sortedPayload)
}
