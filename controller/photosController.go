package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PhotosIndex(c *gin.Context) {
	c.String(http.StatusOK, "There is alot of photos here")
}
