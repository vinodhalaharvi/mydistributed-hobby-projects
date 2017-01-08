# mydistributed-hobby-projects

```
git clone https://github.com/vinodhalaharvi/mydistributed-hobby-projects.git
```
Setup your $GOPATH, and install master and worker processes, usually something
like this 

```
export GOPATH="$RELATIVEPATH/distributed-grep"
alias gopath="cd $GOPATH"
export PATH=$PATH:$GOPATH/bin
```

```
go install github.com/vhalaharvi/distributed-grep/worker
go install github.com/vhalaharvi/distributed-grep/master
```
In one terminal, start workers: 

```
worker
```

In another terminal start master, that searches for a line "testing" in all the
worker machine's /tmp/ directory using 'grep -nr' command
```
master
```

Mockup
All workers are mocked up distributed processes that run on a single box but use different ports for masters to connect as I am only doing this for fun. But it should not take long  to distributed each worker on individual box.
```
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
```



