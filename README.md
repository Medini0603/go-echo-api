# go-echo-api
```bash
#dependencies
go mod init filepath
#get the framework
go get github.com/labstack/echo/v4
go get github.com/stretchr/testify/assert
#run
go run main.go
#run test
go test
go test -v
go test -coverprofile=coverage

```


# APIs written
## No persistent storage
1. get/
1. get/foo
1. get/foo/:id
1. get/foo/sum
1. delete/foo/:id
1. post/foo

## 100% test coverage
