.PHONY: all clean

all:
	go build -o ./server ./main.go
	cd ../proxy-client && go build -o ./client ./main.go

clean:
	rm -rf ./server ../proxy-client/client