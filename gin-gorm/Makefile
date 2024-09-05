.PHONY: init get # phony target

run: build
	./main

init:
	echo "go mod init bronya.com/gin-gorm"

get:
	@echo "==================== get --update ===================="
	go get -u github.com/spf13/viper # 读取配置文件
	go get -u github.com/gin-gonic/gin # HTTP 服务器框架
	go get -u github.com/swaggo/swag/cmd/swag
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files
	go get -u go.uber.org/zap
	@echo "==================== install -o $GOBIN/swag ===================="
	echo "export GOBIN='$HOME/bin'"
	go install github.com/swaggo/swag/cmd/swag@latest

tidy:
	go mod tidy

build:
	go build -o ./main ./src/main.go

clean:
	rm -rf ./main ./log/*

test:
	cd ./test && \
	go test -run TestBuildByMap && \
	go test -run TestBuildByDirNode
	@echo "Test Done!"

swag:
	cd ./src && swag init # 将在 src 目录下创建 docs 目录
