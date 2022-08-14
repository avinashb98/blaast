package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Status struct {
	Message string `json:"message"`
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, Status{Message: "Healthy!"})
}
