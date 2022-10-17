package utils

import (
	"os"
	"time"

	"github.com/robfig/cron"
)

func initCron() {
	c := cron.New()
	registerCrons(c)
	c.Start()
}

func registerCrons(c *cron.Cron) {
	c.AddFunc("0 0 1 * * ?", tmpFileClearCron)
}

func tmpFileClearCron() {
	tmpFileDirPath := CreateSafeFilePath([]string{StaticAssetsPath, "tmp"}, "")
	dirList, _ := os.ReadDir(tmpFileDirPath)
	for _, dir := range dirList {
		dirCreateTime, err := time.Parse("2006-01-02", dir.Name())

		if err != nil {
			os.RemoveAll(tmpFileDirPath + dir.Name())
			Logger.Errorln("assets: 临时文件夹命名错误:", dir.Name()+err.Error())
			return
		}

		if dirCreateTime.Unix() < time.Now().Unix()-60*60*24*7 {
			os.RemoveAll(tmpFileDirPath + dir.Name())
		}
	}
}
