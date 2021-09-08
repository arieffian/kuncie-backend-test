package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/arieffian/kuncie-backend-test/internal/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var (
	fileLogs      *rotatelogs.RotateLogs
	fileLogFormat = "/service.%Y-%m-%d.log"
	fileLogMaxAge = config.GetInt("log.max.age")
	logsPath      string
	err           error
)

// InitLogRotate initialize log rotate
func InitLogRotate() {
	fmt.Println("Starting LOG Rotate")

	//Get the base file dir
	baseDir, _ := os.Getwd()
	fmt.Println("Base directory: ", baseDir)

	//Creating logs directory
	logsPath = filepath.Join(baseDir, config.Get("log.path"))
	os.MkdirAll(logsPath, 0755)
	fmt.Println("Log directory: ", logsPath)

	//Creating file log
	LogFilePath := logsPath + fileLogFormat
	fileLogs, err = rotatelogs.New(LogFilePath,
		rotatelogs.WithMaxAge(time.Duration(fileLogMaxAge)*24*time.Hour))
	if err != nil {
		fmt.Println(err)
	}

}

// GetFileLog get log file
func GetFileLog() *rotatelogs.RotateLogs {
	return fileLogs
}
