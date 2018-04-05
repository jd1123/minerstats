package main

import (
	"fmt"
	"os"

	"bitbucket.org/minerstats/miners/bminer"
	"bitbucket.org/minerstats/miners/ccminer"
	"bitbucket.org/minerstats/miners/claymore"
	"bitbucket.org/minerstats/miners/ethminer"
	"bitbucket.org/minerstats/miners/ewbf"
	"bitbucket.org/minerstats/miners/xmrig"
	"bitbucket.org/minerstats/miners/zm"
	"bitbucket.org/minerstats/sniff"
)

var buf []byte
var host string

//var port string

func usage() {
	var usage string = `minerstats
usage:
	minerstats takes no arguments
	it can detect the following miners:
		bminer 
		ccminer
		claymore
		dtsm
		ewbf
		xmrig-nvidia

	`
	fmt.Println(usage)
}

func hitMiner(miner string, port string) {
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
	case "ethminer":
		{
			ethminer.HitEthminer(host, port, &buf)
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

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		usage()
		os.Exit(1)
	}
	host = "localhost"

	op, err := sniff.SniffMiner()

	// FIXME: return JSON with error
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, v := range op {
		hitMiner(v.Name, v.Port)
	}

}
