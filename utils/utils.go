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

// GetCurrentUserId 获取当前登陆的userId
func GetCurrentUserId(c *gin.Context) (uid string) {
	if userInfo, has := c.Get("claims"); has {
		if claims, ok := ToClaims(userInfo); ok {
			return claims.UserId
		} else {
			return
		}
	} else {
		return
	}
}

// CheckAuthByUserId 通过userId判断当前有无权限
func CheckAuthByUserId(c *gin.Context, userId string) (isAuth bool) {
	currentUid := GetCurrentUserId(c)
	return currentUid == userId
}

// ToHashFileName 转唯一文件名
func ToHashFileName(fileName string) (hashFileName string) {
	if filenames := strings.Split(fileName, "."); len(filenames) < 2 {
		return fmt.Sprintf("%s__%s", fileName, UUID())
	}

	return strings.Join(strings.Split(fileName, "."), fmt.Sprintf("__%s.", UUID()))
}

// If 模拟三元运算符
func If[T any](is bool, v1 T, v2 T) T {
	if is {
		return v1
	}
	return v2
}

func Validate(data any) {

}
