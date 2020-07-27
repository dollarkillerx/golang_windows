# golang_windows
golang windows 编程  

```cassandraql
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o main.dll -buildmode=c-shared main.go
```