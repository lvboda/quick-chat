package utils

import (
	"os"
	"path/filepath"
)

var (
	ConfigFilePath   = ToAbsPath("../config/config.toml")
	CertFilePath     = ToAbsPath("../config/tls.pem")
	KeyFilePath      = ToAbsPath("../config/tls.key")
	StaticAssetsPath = ToAbsPath("../assets")
	LogDirPath       = ToAbsPath("../logs")
)

// GetExecDirPath 获取执行文件外文件夹绝对路径
func GetExecDirPath() string {
	ex, err := os.Executable()
	if err != nil {
		Logger.Fatalln("执行文件路径获取错误: ", err)
	}
	return filepath.Dir(ex)
}

// ToAbsPath 相对路径转绝对路径
func ToAbsPath(filePath string) string {
	return filepath.Join(GetExecDirPath(), filePath)
}
