// Package kafka 往 kafka 写日志模块
package kafka

import (
    "fmt"
    "github.com/Shopify/sarama"
    "time"
)

type logData struct {
    topic string
    data  string
}

var (
    client      sarama.SyncProducer // 声明一个全局连接连接kafka的生产者client
    LogDataChan chan *logData
)

func Init(address []string, maxSize int) (err error) {
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
    // 初始化 logDataChan
    LogDataChan = make(chan *logData, maxSize)

    // 开启goroutine， 后台从通道中取数据发往kafka
    go sendToKafka()
    return
}

// SendToChan 给外部暴露一个函数，该函数把日志数据发送到一个内部的 channel 中
func SendToChan(topic, data string) {
    logd := &logData{
        topic: topic,
        data:  data,
    }
    LogDataChan <- logd
}

// 往 kafka 发送日志的函数
func sendToKafka() {
    for {
        select {
        case logMsg := <-LogDataChan:
            // 构造一个消息
            msg := &sarama.ProducerMessage{}
            msg.Topic = logMsg.topic
            msg.Value = sarama.StringEncoder(logMsg.data)

            // 发送到kafka
            pid, offset, err := client.SendMessage(msg)
            if err != nil {
                fmt.Println("send msg failed, err: ", err)
                return
            }
            fmt.Printf("pid:%v offset:%v\n", pid, offset)
        default:
            time.Sleep(time.Millisecond * 50)
        }

    }
}
