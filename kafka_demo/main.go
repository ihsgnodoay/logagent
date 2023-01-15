package main

import (
    "fmt"
    "github.com/Shopify/sarama"
)

func main() {
    config := sarama.NewConfig()
    // tailf 包使用
    // 发送完数据需要
    config.Producer.RequiredAcks = sarama.WaitForAll
    // leader 和 follow都确认，新选出一个partition
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true

    // 构造一个消息
    msg := &sarama.ProducerMessage{}
    msg.Topic = "web_log"
    msg.Value = sarama.StringEncoder("this is a test log")

    // 连接kafka
    client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
    if err != nil {
        fmt.Println("producer closed, err: ", err)
        return
    }
    fmt.Println("连接kafka成功")
    defer client.Close()

    // 发送数据
    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send msg failed, err: ", err)
        return
    }
    fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
