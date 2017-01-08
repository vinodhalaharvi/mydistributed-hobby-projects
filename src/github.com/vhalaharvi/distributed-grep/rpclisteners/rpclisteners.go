package rpclisteners

import (
    "fmt"
    "os/exec"
    "bytes"
)

type Listener int

type Input struct {
    Line string
    DirPath string
}
/* methods */

func (l *Listener) GrepWorker(input *Input, result *string) error {
    fmt.Printf("Request for Grep: Line: %s File: %s\n", input.Line, input.DirPath)
	cmd := exec.Command("/usr/bin/grep", "-nr", input.Line, input.DirPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
    fmt.Printf("%s\n", err)
    *result = out.String()
	return nil
}

func (l *Listener) GetLine(line []byte, ack *bool) error {
	fmt.Println(string(line))
    *ack = false
	return nil
}
