// Package kafka 往 kafka 写日志模块
package kafka

import (
    "fmt"
    "github.com/Shopify/sarama"
)

var (
    client sarama.SyncProducer // 声明一个全局连接连接kafka的生产者client
)

func Init(address []string) (err error) {
    config := sarama.NewConfig()
    // tailf 包使用 发送完数据需要
    config.Producer.RequiredAcks = sarama.WaitForAll
    // leader 和 follow都确认，新选出一个partition
    config.Producer.Partitioner = sarama.NewRandomPartitioner
    config.Producer.Return.Successes = true

    // 连接kafka
    client, err = sarama.NewSyncProducer(address, config)
    if err != nil {
        fmt.Println("producer closed, err: ", err)
        return
    }
    fmt.Println("连接kafka成功")
    return
}

func SendToKafka(topic, data string) {
    // 构造一个消息
    msg := &sarama.ProducerMessage{}
    msg.Topic = topic
    msg.Value = sarama.StringEncoder(data)

    // 发送到kafka
    pid, offset, err := client.SendMessage(msg)
    if err != nil {
        fmt.Println("send msg failed, err: ", err)
        return
    }
    fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
