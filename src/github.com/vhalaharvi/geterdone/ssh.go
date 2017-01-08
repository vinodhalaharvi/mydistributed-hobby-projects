package main

import (
    "golang.org/x/crypto/ssh"
    "bytes"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "time"
)

func ReadPrivateKeyFile() ([]byte, error)  {
    key, err := ioutil.ReadFile("/Users/vhalaharvi/.ssh/id_rsa")
    if err != nil {
        log.Fatalf("unable to read private key: %v", err)
    }
    return key, err
}



func GetSSHConfig() *ssh.ClientConfig {
    key, err := ReadPrivateKeyFile()
    if err != nil {
        log.Fatalf("unable to read private key: %v", err)
    }

    signer, err := ssh.ParsePrivateKey(key)
    if err != nil {
        log.Fatalf("unable to parse private key: %v", err)
    }

    config := &ssh.ClientConfig{
        User: "vhalaharvi",
        Auth: []ssh.AuthMethod{
            ssh.PublicKeys(signer),
        },
    }
    return config
}


func executeCmd(cmd string, hostname string, config *ssh.ClientConfig) string {
    client, err := ssh.Dial("tcp", hostname, config)
    if err != nil {
        log.Fatal("Failed to dial: ", err)
                }

    session, err := client.NewSession()
    if err != nil {
        log.Fatal("Failed to create session: ", err)
    }
    defer session.Close()

    var b bytes.Buffer
    session.Stdout = &b
    if err := session.Run(cmd); err != nil {
        log.Fatal("Failed to run: " + err.Error())
    }
    return "hostname: " + b.String()
}



func SSHMain() {
    config := GetSSHConfig()
    hosts := [2]string{"localhost:22", "localhost:22"}

    cmd := os.Args[1]
    //hosts := os.Args[2:]

    results := make(chan string, 10)
    timeout := time.After(5 * time.Second)

    for _, hostname := range hosts {
        go func(hostname string) {
            results <- executeCmd(cmd, hostname, config)
        }(hostname)
    }

    for i := 0; i < len(hosts); i++ {
        select {
        case res := <-results:
            fmt.Print(res)
        case <-timeout:
            fmt.Println("Timed out!")
            return
        }
    }
}


