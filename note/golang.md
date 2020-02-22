# go mod
go mod init projectname(工程的根目录)
go list -m  可以查看当前的依赖和版本
go mod edit -fmt 格式化go.mod文件
go mod edit -require=path@version 添加依赖或修改依赖版本
go mod tide 从go.mod删除不需要的依赖，新增需要的依赖，不会改变依赖版本
go mod vendor 生成vendor文件夹
go mod why 
go build -mod=readonly 防止隐士修改go.mod, 如果遇到隐式修改的情况会报错，用来测试go.mod中的依赖是否整洁


## grpc 
protoc 
protoc -I=./ ./helloworld.proto --go_out=plugins=grpc:./
protoc --go_out=plugins=grpc:. *.proto

## go-micro
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
