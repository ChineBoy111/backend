.PHONY: init get # phony target

run: build
	./main

init:
	echo go mod init bronya.com/gin-gorm

get:
	echo go get github.com/spf13/viper
	echo go get -u github.com/gin-gonic/gin

tidy:
	go mod tidy

build:
	go build -o ./main ./main.go

clean:
	rm -rf ./build

test:
	cd ../test && \
	go test -run TestBuildByMap && \
	go test -run TestBuildByDirNode
	@echo Test Done!