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
