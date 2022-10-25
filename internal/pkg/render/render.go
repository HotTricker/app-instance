package render

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	CODE_OK              = 0
	CODE_ERR_APP         = 1001
	CODE_ERR_PARAM       = 1002
	CODE_ERR_DATA_REPEAT = 1003
)

func ParamError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CODE_ERR_PARAM,
		"message": message,
	})
}

func RepeatError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CODE_ERR_DATA_REPEAT,
		"message": message,
	})
}

func AppError(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CODE_ERR_APP,
		"message": message,
	})
}

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    CODE_OK,
		"message": "success",
	})
}
