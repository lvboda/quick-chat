package apiV1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lvboda/quick-chat/utils"
	"github.com/lvboda/quick-chat/utils/status"
)

// Upload 上传文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_FILE_PARSE, err, nil))
		return
	}

	// 文件路径命名规则: /assets/fulltime/年-月-日/fileName__uuid.suffix
	date := time.Now().Format("2006-01-02")
	var fileName string
	if filenames := strings.Split(file.Filename, "."); len(filenames) < 2 {
		fileName = fmt.Sprintf("%s__%s", file.Filename, utils.UUID())
	} else {
		fileName = strings.Join(strings.Split(file.Filename, "."), fmt.Sprintf("__%s.", utils.UUID()))
	}
	src := utils.CreateSafeFilePath([]string{"./assets", "fulltime", date}, fileName)

	err = c.SaveUploadedFile(file, src)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_FILE_UPLOAD, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, strings.TrimPrefix(src, ".")))
}

// Upload 上传临时文件
func UploadTempFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_FILE_PARSE, err, nil))
		return
	}

	// 文件路径命名规则: /assets/temp/年-月-日/fileName__uuid.suffix
	date := time.Now().Format("2006-01-02")
	var fileName string
	if filenames := strings.Split(file.Filename, "."); len(filenames) < 2 {
		fileName = fmt.Sprintf("%s__%s", file.Filename, utils.UUID())
	} else {
		fileName = strings.Join(strings.Split(file.Filename, "."), fmt.Sprintf("__%s.", utils.UUID()))
	}
	src := utils.CreateSafeFilePath([]string{"./assets", "temp", date}, fileName)

	err = c.SaveUploadedFile(file, src)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, status.GetResponse(status.ERROR_FILE_UPLOAD, nil, nil))
		return
	}

	c.JSON(http.StatusOK, status.GetResponse(status.SUCCESS, nil, strings.TrimPrefix(src, ".")))
}
