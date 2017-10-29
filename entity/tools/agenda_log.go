package tools

import(
	model "github.com/txzdream/agenda-go/entity/model"
	"log"
	"os"
	"fmt"
)

const (
	ERROR = "error"
	INFO = "info"
)

// 记录日志--用户名、消息类型（true:info; false:error）、消息内容
// 日志格式：//userName [error/info] message
func LogInfoOrErrorIntoFile(userName string, messageType bool, message string) bool {
	file, err := os.OpenFile(model.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {	// 打开文件失败
		return false
	}

	var messageTypeString string	// 判断消息类型
	if messageType {
		messageTypeString = INFO
	} else {
		messageTypeString = ERROR
	}

	defer file.Close()
	log.SetOutput(file)
	log.Println(fmt.Sprintf("%-15s [%-5s] %s", userName, messageTypeString, message))
	return true
}

