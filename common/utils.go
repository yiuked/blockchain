package common

import (
	"block-chain/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Dispatch(c *gin.Context, code string, desc string, data interface{}) {
	c.JSON(http.StatusOK, models.Response{RespCode: code, RespDesc: desc, RespData: data})
	c.Abort()
}

func StrToInt(int string) int {
	value, err := strconv.Atoi(int)
	if err != nil {
		return 0
	}
	return value
}
