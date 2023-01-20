package main

import (
    "context"
    "fmt"
    clientv3 "go.etcd.io/etcd/client/v3"
    "time"
)

func main() {
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{"127.0.0.1:2379"},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        fmt.Printf("connect to etcd failed, err:%v\n", err)
        return
    }
    fmt.Println("connect to etcd success")
    defer cli.Close()

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    value := `[{"path":"/tmp/nginx.log","topic":"web-log"},{"path":"/tmp/redis.log","topic":"redis-log"}]`
    _, err = cli.Put(ctx, "/etcd_conf_new", value)
    cancel()
    if err != nil {
        fmt.Printf("put to etcd failed, err:%v\n", err)
        return
    }
}
