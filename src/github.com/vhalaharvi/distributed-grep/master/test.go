package main

import (
    "log"
    "net/rpc"
    "time"
    "fmt"
)

type UnitOfWork struct {
    remoteMethod []byte
    RpcClientFunction func(remoteMethod []byte) Result
}

const (
    NUMBER_OF_WORKERS = 2
)

type Result string

func DoWork(workers []*UnitOfWork) {
    c := make(chan Result)
    timeout := time.After(1180 * time.Millisecond)
    go func() { c <- workers[0].RpcClientFunction(workers[0].remoteMethod) }()
    go func() { c <- workers[1].RpcClientFunction(workers[1].remoteMethod) }()

    for {
        select {
        case result := <-c:
            fmt.Printf("%s\n", result)
        case <-timeout:
            fmt.Println("timed out")
            return
        }
    }
    return
}

func MainTest() {
    var reply bool
    workers := []*UnitOfWork{
        &UnitOfWork{
            remoteMethod: []byte("Listener.GetLine"),
            RpcClientFunction: func(remoteMethod []byte) Result  {
                client, err := rpc.Dial("tcp", "localhost:1234")
                if err != nil {
                    log.Fatal(err)
                }
                err = client.Call("Listener.GetLine" , []byte("localhost0"), &reply)
                if err != nil {
                    log.Fatal(err)
                }
                return "done"
            } },
        &UnitOfWork{
            remoteMethod: []byte("Listener.GetLine"),
            RpcClientFunction: func(remoteMethod []byte) Result {
                client, err := rpc.Dial("tcp", "localhost:1234")
                if err != nil {
                    log.Fatal(err)
                }
                err = client.Call("Listener.GetLine", []byte("localhost1"), &reply)
                if err != nil {
                    log.Fatal(err)
                }
                return "done"
            } } }
    DoWork(workers)
}
