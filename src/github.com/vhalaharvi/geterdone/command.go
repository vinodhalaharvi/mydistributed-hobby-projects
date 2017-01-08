package main

import (
    "fmt"
    "io"
    "log"
    "os/exec"
    "bytes"
)

func LookPath() {
    path, err := exec.LookPath("fortune")
    if err != nil {
        log.Fatal("installing fortune is in your future")
    }
    fmt.Printf("fortune is available at %s\n", path)
}

func CommandStdinStdout(input []byte, args...string) io.Writer{
    cmd := exec.Command(args[0], args[1:]...)
    cmd.Stdin = bytes.NewReader(input)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return cmd.Stdout
}

func CommandOutput(args...string) []byte{
    out, err := exec.Command(args[0], args[1:]...).Output()
    if err != nil {
        log.Fatal(err)
    }
    return out
}

func Command(args...string) {
    cmd := exec.Command(args[0], args[1:]...)
    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Waiting for command to finish...")
    err = cmd.Wait()
    log.Printf("Command finished with error: %v", err)
}

func CommandStdoutPipe(args...string) io.ReadCloser{
    cmd := exec.Command(args[0], args[1:]...)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
    return stdout
}

func CommandStdinPipeStdout(input []byte, args...string) []byte{
    cmd := exec.Command(args[0], args[1:]...)
    stdin, err := cmd.StdinPipe()
    if err != nil {
        log.Fatal(err)
    }
    go func() {
        defer stdin.Close()
        stdin.Write(input)
    }()
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    return out
}

func CommandStdinPipeStdoutPipeStderrPipe(input []byte, args...string) ([]byte, error){
    cmd := exec.Command(args[0], args[1:]...)
    stdin, err := cmd.StdinPipe()
    if err != nil {
        log.Fatal(err)
    }
    go func() {
        defer stdin.Close()
        stdin.Write(input)
    }()
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    return out, err
}

func CommandStderrPipe(args...string) io.ReadCloser {
    cmd := exec.Command(args[0], args[1:]...)
    stderr, err := cmd.StderrPipe()
    if err != nil {
        log.Fatal(err)
    }
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }
    //slurp, _ := ioutil.ReadAll(stderr)
    //fmt.Printf("%s\n", slurp)
    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }
    return stderr
}

func CommandCombinedOutput(args...string) []byte{
    cmd := exec.Command(args[0], args[1:]...)
    stdoutStderr, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatal(err)
    }
    return stdoutStderr
}
