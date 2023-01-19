package taillog

import "github.com/dyihs/logagent/etcd"

var tskMgr *taillogMgr

type taillogMgr struct {
    logEntry []*etcd.LogEntry
    // tskMap   map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry) {
    // 根据路径打开对应的日志
    tskMgr = &taillogMgr{
        logEntry: logEntryConf,
    }
    for _, logEntry := range logEntryConf {
        NewTailTask(logEntry.Path, logEntry.Topic)
    }
}
