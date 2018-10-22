package main

import (
	"fmt"
	"os"

	"github.com/jd1123/minerstats/miners/bminer"
	"github.com/jd1123/minerstats/miners/ccminer"
	"github.com/jd1123/minerstats/miners/claymore"
	"github.com/jd1123/minerstats/miners/ethminer"
	"github.com/jd1123/minerstats/miners/ewbf"
	"github.com/jd1123/minerstats/miners/xmrig"
	"github.com/jd1123/minerstats/miners/zm"
	"github.com/jd1123/minerstats/output"
	"github.com/jd1123/minerstats/sniff"
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
	host = "localhost"
  if (len(args) == 0) {
    op, err := sniff.SniffMiner()

    // FIXME: return JSON with error
    if err != nil {
      //		fmt.Println(err.Error())
      fmt.Println(string(output.MakeJSONError("", err)))
      os.Exit(1)
    }

    for _, v := range op {
      hitMiner(v.Name, v.Port)
    }
  } else {
    if (args[0] == "-h") {
      usage()
      os.Exit(1)
    } else {
      hitMiner(args[0], args[1])
    }
  }
}
