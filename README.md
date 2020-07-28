# golang_windows
golang windows 编程  

```cassandraql
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o main.dll -buildmode=c-shared main.go
```


参考:
- https://www.cnblogs.com/Kingram/p/12088087.html
- https://blog.csdn.net/rtduq/article/details/80340461
- https://blog.csdn.net/wn0112/article/details/100708412~