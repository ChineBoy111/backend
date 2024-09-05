.PHONY: init get # phony target

run: build
	./main

init:
	echo go mod init bronya.com/gin-gorm

get:
	go get github.com/spf13/viper
	go get -u github.com/gin-gonic/gin
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest
	go get -u github.com/swaggo/gin-swagger
	go get -u github.com/swaggo/files

install:


tidy:
	go mod tidy

build:
	go build -v -o ./main ./main.go

clean:
	rm -rf ./build

test:
	cd ../test && \
	go test -run TestBuildByMap && \
	go test -run TestBuildByDirNode
	@echo Test Done!

swag:
	swag -v
	swag init # 将在 src 目录下创建 docs 目录
