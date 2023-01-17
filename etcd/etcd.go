package etcd

import (
    "fmt"
    clientv3 "go.etcd.io/etcd/client/v3"
    "time"
)

// Init etcd 初始化
func Init(addr string, timeout time.Duration) (err error) {
    _, err = clientv3.New(clientv3.Config{
        Endpoints:   []string{"127.0.0.1:2379"},
        DialTimeout: 5 * time.Second,
    })

    if err != nil {
        fmt.Printf("connect to etcd failed, err:%v\n", err)
        return
    }
    return
}
