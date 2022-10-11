package v1

import (
	"net/http"

	"github.com/easonnong/grade-manage/model"
	"github.com/easonnong/grade-manage/utils/errmsg"
	"github.com/gin-gonic/gin"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url, code := model.UpLoadFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"url":     url,
	})
}
