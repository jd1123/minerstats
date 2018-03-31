default:
	make all

build:
	go build -ldflags="-s -w"

clean:
	rm -rf minerstats

pack:
	upx --brute minerstats

all:
	make clean
	make build
#	make pack
