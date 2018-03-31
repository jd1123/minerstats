package ewbf

import (
	"bytes"
	"encoding/json"

	"bitbucket.org/minerstats/dialminer"
	"bitbucket.org/minerstats/output"
)

type Result struct {
	Gpuid          int
	Cudaid         int
	Busid          string
	Name           string
	GpuStatus      int
	Solver         int
	Temp           int
	GPUPower       int
	Hashrate       int `json:"speed_sps"`
	AcceptedShares int
	RejectedShares int
	StartTime      int
}

type EWBFOut struct {
	Id               int      `json:"id"`
	Method           string   `json:"method"`
	Error            string   `json:"error"`
	StartTime        int      `json:"start_time"`
	CurrentServer    string   `json:"current_server"`
	AvailableServers int      `json:"available_servers"`
	ServerStatus     int      `json:"server_status"`
	Results          []Result `json:"result"`
}

func parseOutput(b []byte) *EWBFOut {
	e := new(EWBFOut)
	e.Results = make([]Result, 1, 30)
	json.Unmarshal(b, &e)
	return e
}

func HitEwbf(host_l string, port_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0

	bu, err := dialminer.DialMiner(host_l, port_l, "{\"method\":\"getstat\"}\n\n")
	if err != nil {
		panic(err)
	}

	bu = bytes.Trim(bu, "\x00")
	e := parseOutput(bu)

	for _, v := range e.Results {
		hrtotal += float64(v.Hashrate)
		numMiners++
	}

	o := output.NewOutput()
	o.Minername = "ewbf"
	o.Hashrate = hrtotal
	o.NumMiners = numMiners
	js, _ := json.Marshal(o)
	*buf = js
}
