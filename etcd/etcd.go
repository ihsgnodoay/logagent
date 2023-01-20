package etcd

import (
    "context"
    "encoding/json"
    "fmt"
    clientv3 "go.etcd.io/etcd/client/v3"
    "time"
)

var (
    cli *clientv3.Client
)

type LogEntry struct {
    Path  string `json:"path"`
    Topic string `json:"topic"`
}

// Init etcd 初始化
func Init(addr string, timeout time.Duration) (err error) {
    cli, err = clientv3.New(clientv3.Config{
        Endpoints:   []string{"127.0.0.1:2379"},
        DialTimeout: 5 * time.Second,
    })

    if err != nil {
        fmt.Printf("connect to etcd failed, err:%v\n", err)
        return
    }
    return
}

// GetConf 根据 key 从etcd中获取配置项
func GetConf(key string) (logEntryCOnf []*LogEntry, err error) {
    // get
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    resp, err := cli.Get(ctx, key)
    cancel()
    if err != nil {
        fmt.Printf("get from etcd failed, err:%v\n", err)
        return
    }

    for _, ev := range resp.Kvs {
        err = json.Unmarshal(ev.Value, &logEntryCOnf)
        if err != nil {
            fmt.Printf("Unmarshal etcd conf failed, err:%v\n", err)
            return
        }
        fmt.Printf("%s:%s\n", ev.Key, ev.Value)
    }
    return
}

// WatchConf watch conf
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
    ch := cli.Watch(context.Background(), key)
    // 从通道中尝试获取值
    for wresp := range ch {
        for _, evt := range wresp.Events {
            fmt.Printf("Type: %v\n key:%v value:%v\n", evt.Type, string(evt.Kv.Value), string(evt.Kv.Value))
            // 通知 taillog.tskMgr
            var newConf []*LogEntry
            err := json.Unmarshal(evt.Kv.Value, &newConf)
            if err != nil {
                fmt.Printf("unmarshal failed, err:%v\n", err)
                continue
            }
            fmt.Printf("get new conf:%v\n", err)
            newConfCh <- newConf
        }
    }
}
