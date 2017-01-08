package main

import (
    "sync"
    "fmt"
    "log"
    "net"
    "net/rpc"
    "github.com/vhalaharvi/distributed-grep/rpclisteners"
)


func MainWorker() {

    listenAddrs := []string{
        "0.0.0.0:1234",
        "0.0.0.0:1235",
        "0.0.0.0:1236",
        "0.0.0.0:1237",
        "0.0.0.0:1238",
        "0.0.0.0:1239",
        "0.0.0.0:1240",
        "0.0.0.0:1241",
        "0.0.0.0:1242",
    }
    var wg sync.WaitGroup

    wg.Add(len(listenAddrs))
    for _, listenAddr := range listenAddrs {
        fmt.Printf("%s\n", listenAddr)
        addr, err := net.ResolveTCPAddr("tcp", listenAddr)
        if err != nil {
            log.Fatal(err)
        }
        inbound, err := net.ListenTCP("tcp", addr)
        if err != nil {
            log.Fatal(err)
        }
        listener := new(rpclisteners.Listener)
        rpc.Register(listener)
        go func() {
            for {
                rpc.Accept(inbound)
            }
            wg.Done()
        }()
    }
    wg.Wait()
}
