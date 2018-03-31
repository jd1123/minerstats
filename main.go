package main

import (
	"fmt"
	"os"

	"bitbucket.org/minerstats/bminer"
	"bitbucket.org/minerstats/ccminer"
	"bitbucket.org/minerstats/claymore"
	"bitbucket.org/minerstats/ewbf"
	"bitbucket.org/minerstats/zm"
)

var buf []byte
var host string
var port string

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
			ccminer.HitCCMiner(host, port, &buf)
		}
	case "bminer":
		{
			bminer.HitBminer(host, port, &buf)
		}
	case "ewbf":
		{
			ewbf.HitEwbf(host, port, &buf)
		}
	case "zm":
		{
			zm.HitZM(host, port, &buf)
		}
	case "claymore":
		{
			claymore.HitClaymore(host, port, &buf)
		}
	default:
		{
			fmt.Println("ERROR! miner argument not recognized!")
			os.Exit(1)
		}
	}
	fmt.Println(string(buf))
}
