package main

import (
	"encoding/json"
	"fmt"
	"ginchat/utils/logger"
)

type logmessage struct {
	message string
	code    int
}

func main() {

	logger.SetOutputPath("/project/logs/logfile.log")
	count := 100

	for i := 0; i < count; i++ {
		logMsg := make(map[int]interface{})
		logMsg[i] = "这是一条日志消息"
		strObj, _ := json.Marshal(logMsg)
		strMsg := string(strObj)
		fmt.Println(strMsg)
		logger.Infof("messge:%s", strMsg)

	}
}
