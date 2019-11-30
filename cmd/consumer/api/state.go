package api

import (
	"github.com/felixvo/lmax/cmd/consumer/state"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewCurrentStateHandler(st *state.State) gin.HandlerFunc{
	return func(c *gin.Context) {
		c.JSON(http.StatusOK,st)
	}
}
