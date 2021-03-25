package initialize

import (
	"frozen-go-project/app/file_server/config"
	"frozen-go-project/app/file_server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.StaticFS(config.UPLOAD_PATH, http.Dir(config.UPLOAD_PATH))
	group := router.Group("")
	initFileUploadAndDownload(group)
	return router
}

func initFileUploadAndDownload(router *gin.RouterGroup) {
	group := router.Group("file")
	group.POST("/upload", uploadFunc)
}

func uploadFunc(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	code, message := 0, "success"
	if err != nil {
		code = 1001
		message = err.Error()
	}
	filepath, filename, err := service.UploadFile(header)
	if err != nil {
		code = 1001
		message = err.Error()
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,
		"message": message,
		"data": map[string]interface{}{
			"filepath": filepath,
			"filename": filename,
		},
	})
}
