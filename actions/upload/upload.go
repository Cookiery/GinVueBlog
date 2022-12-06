package upload

import (
	"main/commond/errmsg"
	"main/model/service/qiniu"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url, err := qiniu.UploadFile(file, fileSize)
	c.JSON(http.StatusOK, gin.H{
		"status":  err,
		"message": errmsg.ErrMsg(err),
		"url":     url,
	})
}
