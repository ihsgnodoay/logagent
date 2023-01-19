// Package taillog 从日志文件收集日志模块
package taillog

import (
    "fmt"
    "github.com/dyihs/logagent/kafka"
    "github.com/hpcloud/tail"
)

type TailTask struct {
    path     string
    topic    string
    instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
    tailObj = &TailTask{
        path:  path,
        topic: topic,
    }
    tailObj.init()
    return
}

// 根据路径打开对应的日志
func (t *TailTask) init() {
    c := tail.Config{
        ReOpen:    true,                                 // 重新打开
        Follow:    true,                                 // 是否跟随
        Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的那个地方开始读
        MustExist: false,                                // 文件不存在不报错
        Poll:      true,
    }
    var err error
    t.instance, err = tail.TailFile(t.path, c)
    if err != nil {
        fmt.Printf("tail file failed, err%v\n:", err)
        return
    }
    go t.run()
}

// 采集日志发送到 kafka
func (t *TailTask) run() {
    for {
        select {
        // 从tailObj通道中一行一行读取日志
        case lines := <-t.instance.Lines:
            // 发往 kafka
            // kafka.SendToKafka(t.topic, lines.Text)
            // 先把日志数据发到一个通道中，kafka 包中有单独的 goroutine 去取日志数据发到 kafka
            kafka.SendToChan(t.path, lines.Text)

        }
    }
}
