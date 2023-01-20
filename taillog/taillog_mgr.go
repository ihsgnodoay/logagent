package taillog

import (
    "fmt"
    "github.com/dyihs/logagent/etcd"
    "time"
)

var tskMgr *taillogMgr

// tailTask 管理者
type taillogMgr struct {
    logEntry    []*etcd.LogEntry
    taskMap     map[string]*TailTask
    newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
    // 根据路径打开对应的日志
    tskMgr = &taillogMgr{
        logEntry:    logEntryConf,
        taskMap:     make(map[string]*TailTask, 16),
        newConfChan: make(chan []*etcd.LogEntry), // 无缓冲区的通道
    }
    for _, logEntry := range logEntryConf {
        NewTailTask(logEntry.Path, logEntry.Topic)
    }
    go tskMgr.run()
}

// 监听newConfChan，有了新的配置过来之后就做对应的处理
// 配置新增、删除、更改
func (t *taillogMgr) run() {
    for {
        select {
        case newConf := <-t.newConfChan:
            fmt.Println("新的配置来了：", newConf)
        default:
            time.Sleep(time.Second)
        }
    }
}

// NewConfChan 一个函数，向外暴露tskMgr的newConfChan
func NewConfChan() chan<- []*etcd.LogEntry {
    return tskMgr.newConfChan
}
