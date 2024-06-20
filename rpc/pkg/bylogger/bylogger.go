package bylogger

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"os"
	"path/filepath"
	"runtime"
)

var logger log.Logger

func init() {
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "datetime", log.DefaultTimestampUTC)
	logger = log.With(logger, "application", filepath.Base(os.Args[0]))
}
func makeMessage(inputs []interface{}) string {
	if inputs == nil || len(inputs) == 0 {
		return "(MISSING)"
	}
	var message string
	for _, x := range inputs {
		if message != "" {
			message += " "
		}
		message += fmt.Sprint(x)
	}
	return message
}

func getPathFile(file string) string {
	dir, _ := os.Getwd()
	filePath, err := filepath.Rel(dir, file)
	if err != nil {
		panic(err)
	}
	return filePath
}
func LogInfo(inputs ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	logger.Log("type", "info", "file", fmt.Sprintf("%s:%v", file, line), "message", makeMessage(inputs))
}
func LogWarn(inputs ...interface{}) {
	message := makeMessage(inputs)
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	logger.Log("type", "warning", "file", fmt.Sprintf("%s:%v", file, line), "message", message)
}
func LogErr(inputs ...interface{}) {
	message := makeMessage(inputs)
	_, file, line, _ := runtime.Caller(1)
	file = getPathFile(file)
	logger.Log("type", "error", "file", fmt.Sprintf("%s:%v", file, line), "message", message)
}
