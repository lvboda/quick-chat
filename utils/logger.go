package utils

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func initLogger() {
	Logger = logrus.New()

	logFilePath := CreateSafeFilePath([]string{LogDirPath}, "quick_chat_server.log")

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		Logger.Fatalln("log文件路径解析错误: ", err)
	}

	writers := []io.Writer{file, os.Stdout}

	fileAndStdoutWriter := io.MultiWriter(writers...)

	gin.DefaultErrorWriter = fileAndStdoutWriter
	Logger.SetOutput(fileAndStdoutWriter)
	Logger.SetReportCaller(true)
	Logger.SetLevel(logrus.DebugLevel)
}
