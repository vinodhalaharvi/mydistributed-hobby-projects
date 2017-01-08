package main

import (
    "os"
    "fmt"
    "strconv"
    "log"
    "os/exec"
)

type Kafka struct {}

func (k *Kafka) PreferredReplicaElection(topicName string, zooKeeper string) {
    //replica-election
    app := "kafka-preferred-replica-election.sh "
    argString := app + " --zookeeper " + zooKeeper
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) ReassignPartitionsRollback(topicName string, zooKeeper string, topicsToMoveJson string, brokerList string) {
    //execute rollback
    app := "kafka-reassign-partitions.sh "
    argString := app + " --zookeeper " +
    zooKeeper + " --topics-to-move " +
    topicsToMoveJson + " --broker-list " +
    brokerList + " --execute"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) ReassignPartitionsExecute(topicName string, zooKeeper string, topicsToReassignJson string, brokerList string) {
    //execute
    app := "kafka-reassign-partitions.sh "
    argString := app + " --zookeeper " + zooKeeper +
    " --reassignment-json-file " +  topicsToReassignJson +
    " --broker-list " + brokerList +
    " --execute"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) ReassignPartitionsGenerate(topicName string, zooKeeper string, topicsToMoveJson string, brokerList string){
    //generate reassigned
    //| head -n 6 | tail -1 | python -m json.tool > /tmp/topics-reassigned.json
    // | head -n 3 | tail -1 | python -m json.tool >  /tmp/topics-rollback.json
    app := "kafka-reassign-partitions.sh "
    argString := app +
    " --zookeeper " + zooKeeper +
    " --topics-to-move " +  topicsToMoveJson +
    " --broker-list " + brokerList  +
    " --generate"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) ReassignPartitionsVerify(topicName string, zooKeeper string, topicsToReassignJson string, brokerList string){
    //verify
    app := "kafka-reassign-partitions.sh "
    argString := app +
    " --zookeeper " + zooKeeper +
    " --reassignment-json-file " + topicsToReassignJson +
    " --broker-list " + brokerList +
    " --verify"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}


func (k *Kafka) List(zooKeeper string) {
    app := "kafka-topics.sh "
    argString := " --zookeeper " + zooKeeper +  " --list"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) Create(zooKeeper string, topicName string, partitions int, replicationFactor int){
    //create
    app := "kafka-topics.sh "
    argString := app + " --zookeeper " +
    zooKeeper +
    " --create " +
    " --topic " + topicName +
    " --partitions " + strconv.Itoa(partitions) +
    " --replication-factor " + strconv.Itoa(replicationFactor)
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) Delete(topicName string)  {
}

func (k *Kafka) Describe(topicName string, zooKeeper string)  {
    app := "kafka-topics.sh "
    argString := " --zookeeper " + zooKeeper +  " --describe"
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) Consume(topicName string, zooKeeper string){
    //consumer
    app := "kafka-console-consumer.sh "
    argString := app +
    " --zookeeper " + zooKeeper +
    " --topic " + topicName
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}

func (k *Kafka) Produce(topicName string, zooKeeper string){
    app := "kafka-console-producer.sh "
    argString := app +
    " --zookeeper " + zooKeeper +
    " --topic " + topicName
    fmt.Printf("%s%s\n", app, argString)
    CommandCombinedOutput(app, argString)
}


func SimpleCommand(app string, args ...string) {
    cmd := exec.Command(app, args...)
    stdoutStderr, err := cmd.CombinedOutput()
    fmt.Printf("%s\n", stdoutStderr)
    if err != nil {
        log.Fatal(err)
    }
}

const (
    //ZOOKEEPER = "10.200.4.126:2181,10.200.4.127:2181,10.200.4.128:2181"
    ZOOKEEPER = "172.28.128.21:2181"
)


func InvokeKafka(app string
        , context string
        , action string
        , context string) error{
    k := Kafka{}
    switch action {
    case "list":
        k.List(ZOOKEEPER)
    case "preferred-replica-election":
        k.PreferredReplicaElection(topicName, ZOOKEEPER)
    case "reassign-partition-generate":
        k.ReassignPartitionsGenerate(ZOOKEEPER, "testing", 10, 1)
    case "reassign-partition-verify":
        k.ReassignPartitionsVerify(ZOOKEEPER, "testing", 10, 1)
    case "reassign-partition-rollback":
        k.ReassignPartitionsRollback(ZOOKEEPER, "testing", 10, 1)
    case "delete":
        k.Delete(ZOOKEEPER, "testing", 10, 1)
    case "describe":
        k.Describe(ZOOKEEPER, "testing", 10, 1)
    case "consume":
        k.Consume(ZOOKEEPER, "testing", 10, 1)
    case "produce":
        k.Produce(ZOOKEEPER, "testing", 10, 1)
    default:
        log.Fatal(" Invalid Action found: ", action)
    }
}


//geterdone kafka lambda list topics
func main() {
    //app, context, action, object :=
    app, context, action, object :=
        os.Args[1], os.Args[2], os.Args[3], os.Args[4]
    InvokeKafka(app, context, action, object)
}
