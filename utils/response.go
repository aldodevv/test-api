package utils

import (
	"test-api/index/advanced/constants"
	"test-api/index/advanced/structs"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess membungkus JSON khusus untuk format sukses
// Kode RC otomatis SC00 dari constants untuk semua jenis success.
func ResponseSuccess(c *gin.Context, statusCode int, msg string, data interface{}) {
	c.JSON(statusCode, structs.BaseResponse{
		RC:   constants.RCSukses,
		Msg:  msg,
		Data: data,
	})
}

// ResponseError membungkus JSON khusus untuk error
// Data akan diset nil / null, dan kode RC dinamis berdasarkan fitur yang error
func ResponseError(c *gin.Context, statusCode int, rc string, msg string) {
	c.JSON(statusCode, structs.BaseResponse{
		RC:   rc,
		Msg:  msg,
		Data: nil,
	})
}
