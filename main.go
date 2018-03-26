package main

import (
	"fmt"
	"os"
)

var buf []byte
var host string;
var port string;

func usage() {
	var usage string = `
minerstats
usage:
	minerstats <host> <port> <minerprogram>
	
	host is your miner (localhost)
	port is the port on which the miner api is listening
	miner program can be either "bminer" or "ccminer"
	`
	fmt.Println(usage)
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		usage()
		os.Exit(1)
	} else {
		port = args[1]
		host = args[0]
	}
  var miner = args[2]
  switch miner {
  case "ccminer":
    {
      hitCCMiner()
    }
  case "bminer":
    {
      hitBminer()
    }
  default:
    {
      fmt.Println("ERROR! miner argument not recognized!")
      os.Exit(1)
    }
  }
  fmt.Println(string(buf))
}
