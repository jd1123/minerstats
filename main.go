package main

import (
	"fmt"
	"os"

	"bitbucket.org/minerstats/miners/bminer"
	"bitbucket.org/minerstats/miners/ccminer"
	"bitbucket.org/minerstats/miners/claymore"
	"bitbucket.org/minerstats/miners/ewbf"
	"bitbucket.org/minerstats/miners/xmrig"
	"bitbucket.org/minerstats/miners/zm"
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
	miners supported:
		bminer: "bminer"
		ccminer: "ccminer"
		claymore: "claymore"
		dtsm : "zm"
		ewbf: "ewbf"
		xmrig-nvidia: "xmrig"

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
	case "xmrig":
		{
			xmrig.HitXMRig(host, port, &buf)
		}
	default:
		{
			fmt.Println("ERROR! miner argument not recognized!")
			os.Exit(1)
		}
	}
	fmt.Println(string(buf))
}
