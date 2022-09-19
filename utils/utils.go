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

// 创建安全的文件路径
// 没有文件夹则创建
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

func Validate(data any) {

}
