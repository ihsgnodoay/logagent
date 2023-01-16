// Package taillog 从日志文件收集日志模块
package taillog

import (
    "fmt"
    "github.com/hpcloud/tail"
)

var (
    tailObj *tail.Tail
    LogChan chan string
)

func Init(fileName string) (err error) {
    config := tail.Config{
        ReOpen:    true,                                 // 重新打开
        Follow:    true,                                 // 是否跟随
        Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个地方开始读
        MustExist: false,                                // 文件不存在不报错
        Poll:      true,
    }

    tailObj, err = tail.TailFile(fileName, config)
    if err != nil {
        fmt.Println("tail file failed, err:", err)
        return
    }
    return
}

func ReadChan() <-chan *tail.Line {
    return tailObj.Lines
}
