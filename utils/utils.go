package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func init() {
	initLogger()
	initConfig()
	initDB()
}

// createApp
func CreateApp() *gin.Engine {
	// 设置mode
	mode := GetConfig().Server.Mode
	gin.SetMode(mode)

	app := gin.New()

	return app
}

// CreateSafeFilePath 创建安全的文件路径
func CreateSafeFilePath(dirNames []string, fileName string) string {
	if len(dirNames) < 1 {
		return ""
	}

	var basePath string

	for index, dir := range dirNames {
		if index == 0 {
			basePath = dir
		} else {
			basePath = fmt.Sprintf("%s/%s", basePath, dir)
		}

		if _, err := os.Stat(basePath); os.IsNotExist(err) {
			os.MkdirAll(basePath, os.ModePerm)
			os.Chmod(basePath, os.ModePerm)
		}
	}

	return strings.Join([]string{basePath, fileName}, "/")
}

// MergeJson 合并json
func MergeJson(args ...any) (map[string]any, error) {
	var jsonMap map[string]any

	for _, item := range args {
		itemJson, err := json.Marshal(item)

		if err != nil {
			return nil, err
		}

		json.Unmarshal(itemJson, &jsonMap)
	}

	return jsonMap, nil
}

// UUID 生成唯一id
func UUID() string {
	uuid := uuid.New().String()
	return strings.ReplaceAll(uuid, "-", "")
}

// CheckAuthByUserId 通过userId判断当前有无权限
func CheckAuthByUserId(c *gin.Context, userId string) (isAuth bool) {
	if userInfo, has := c.Get("claims"); has {
		if v, ok := userInfo.(*claims); ok {
			return v.UserId == userId
		} else {
			return
		}
	} else {
		return
	}
}

func Validate(data any) {

}
