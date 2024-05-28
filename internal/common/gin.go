package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(ctx *gin.Context, code int, err interface{}) {
	if code == -1 {
		code = http.StatusInternalServerError
	}

	if str, ok := err.(string); ok {
		ctx.JSON(code, gin.H{
			"error": map[string]string{
				"message": str,
			},
		})
		return
	}

	if e, ok := err.(error); ok {
		ctx.JSON(code, gin.H{
			"error": map[string]string{
				"message": e.Error(),
			},
		})
		return
	}

	ctx.JSON(code, gin.H{
		"error": map[string]string{
			"message": fmt.Sprintf("%v", err),
		},
	})
}
