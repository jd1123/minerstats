package ewbf

import (
	"bytes"
	"encoding/json"
	"fmt"

	"bitbucket.org/minerstats/dialminer"
	"bitbucket.org/minerstats/output"
)

type result struct {
	Gpuid          int
	Cudaid         int
	Busid          string
	Name           string
	GpuStatus      int
	Solver         int
	Temp           int
	GPUPower       int `json:"gpu_power_usage"`
	Hashrate       int `json:"speed_sps"`
	AcceptedShares int
	RejectedShares int
	StartTime      int
}

type ewbfOut struct {
	Id               int      `json:"id"`
	Method           string   `json:"method"`
	Error            string   `json:"error"`
	StartTime        int      `json:"start_time"`
	CurrentServer    string   `json:"current_server"`
	AvailableServers int      `json:"available_servers"`
	ServerStatus     int      `json:"server_status"`
	Results          []result `json:"result"`
}

func parseOutput(b []byte) *ewbfOut {
	e := new(ewbfOut)
	e.Results = make([]result, 1, 30)
	json.Unmarshal(b, &e)
	return e
}

func HitEwbf(host_l string, port_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	var totalPower float64 = 0

	bu, err := dialminer.DialMiner(host_l, port_l, "{\"method\":\"getstat\"}\n\n")
	if err != nil {
		panic(err)
	}

	bu = bytes.Trim(bu, "\x00")
	e := parseOutput(bu)
	for _, v := range e.Results {
		hrtotal += float64(v.Hashrate)
		totalPower += float64(v.GPUPower)
		numMiners++
	}

	js, err := output.MakeJSON_full("ewbf", hrtotal, numMiners, totalPower)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	*buf = js
}
