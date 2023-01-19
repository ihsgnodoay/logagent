// logAgent 入口程序
package main

import (
    "fmt"
    "github.com/dyihs/logagent/config"
    "github.com/dyihs/logagent/etcd"
    "github.com/dyihs/logagent/kafka"
    "github.com/dyihs/logagent/taillog"
    "gopkg.in/ini.v1"
    "time"
)

// var (
//     cfg, err = ini.Load("./config/config.ini")
// )
var (
    cfg = new(config.AppConf)
)

func main() {
    // 0. 加载配置文件
    err := ini.MapTo(cfg, "./config/config.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v\n", err)
        return
    }

    // 1. 初始化kafka连接
    err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
    if err != nil {
        fmt.Printf("init kafka failed：err%v\n", err)
        return
    }
    fmt.Println("init kafka success.")

    // 2. 初始化etcd
    err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
    if err != nil {
        fmt.Printf("init etcd failed：err%v\n", err)
        return
    }
    fmt.Println("init etcd success.")

    // 从 etcd 中获取日志收集项的配置信息
    logEntryConf, err := etcd.GetConf("/etcd_conf")
    if err != nil {
        fmt.Printf("etcd.GetCOnf failed, err: %v\n", err)
        return
    }
    fmt.Println("etcd.GetConf success", logEntryConf)
    for index, value := range logEntryConf {
        fmt.Printf("index: %v, value: %v\n", index, value)
    }

    // 3. 收集日志发往kafka
    taillog.Init(logEntryConf)

}
