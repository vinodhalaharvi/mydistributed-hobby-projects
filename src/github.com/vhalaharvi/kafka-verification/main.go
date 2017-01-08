package main

import (
    "strings"
    "fmt"
    "os"
    "github.com/Shopify/sarama"
)

func SyncProduce() {

    config := sarama.NewConfig()
    config.Producer.RequiredAcks = sarama.WaitForAll
    config.Producer.Retry.Max = 5
    config.Producer.Return.Successes = true

    if os.Getenv("BROKER_LIST_WITH_PORTS") == "" {
        panic("Missing BROKER_LIST_WITH_PORTS. set environment variable BROKER_LIST_WITH_PORTS..")
    }
    brokers := strings.Split(os.Getenv("BROKER_LIST_WITH_PORTS"), ",")
    producer, err := sarama.NewSyncProducer(brokers, config)
    if err != nil {
        // Should not reach here
        panic(err)
    }

    defer func() {
        if err := producer.Close(); err != nil {
            // Should not reach here
            panic(err)
        }
    }()

    topic := "important"
    msg := &sarama.ProducerMessage{
        Topic: topic,
        Value: sarama.StringEncoder("Something Cool"),
    }

    partition, offset, err := producer.SendMessage(msg)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}


func ExampleBroker() {
    broker := sarama.NewBroker(os.Getenv("BROKER_LIST_WITH_PORTS"))
    err := broker.Open(nil)
    if err != nil {
        panic(err)
    }
    request := sarama.MetadataRequest{}
    response, err := broker.GetMetadata(&request)
    if err != nil {
        panic(err)
    }
    for _, topic := range response.Topics {
        fmt.Println(topic)
    }
    if err = broker.Close(); err !=nil {
        panic(err)
    }
}



func main() {
    ExampleBroker()
}
