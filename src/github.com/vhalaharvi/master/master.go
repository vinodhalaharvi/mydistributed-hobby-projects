package main

import (
    "time"
    "fmt"
    "log"
    "net/rpc"
    "sync"
)

type Input struct {
    Line string
    DirPath string
}


func GrepMaster(line string, dirPath string){
    hosts := []string{
        "localhost:1234",
        "localhost:1235",
        "localhost:1236",
        "localhost:1237",
        "localhost:1238",
        "localhost:1239",
        "localhost:1240",
        "localhost:1241",
        "localhost:1242",
    }
    input := Input{
        Line: line,
        DirPath: dirPath,
    }

    var result string
    var wg sync.WaitGroup
    wg.Add(len(hosts))

    c := make(chan string)

    for _, host := range hosts {
        go func() {
            client, err := rpc.Dial("tcp", host)
            if err != nil {
                log.Fatal(err)
            }
            err = client.Call("Listener.GrepWorker" , &input, &result)
            if err != nil {
                log.Fatal(err)
            }
            c <- result
            wg.Done()
        }()
    }

    timeout := time.After(time.Second * 10)
    for i := 0; i < len(hosts); i++ {
        select {
        case result := <-c:
            fmt.Printf("%s\n", result)
        case <-timeout:
            fmt.Printf("%s\n", "time out.")
            return
        }
    }
}
