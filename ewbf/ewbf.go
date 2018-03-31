package ewbf

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/minerstats/output"

	curl "github.com/andelf/go-curl"
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
	var bu []byte
	easy := curl.EasyInit()
	defer easy.Cleanup()
	var hrtotal float64 = 0
	var numMiners int = 0

	easy.Setopt(curl.OPT_URL, "http://shitcoin5:3333")

	// make a callback function
	fooTest := func(b []byte, userdata interface{}) bool {
		//		println("DEBUG: size=>", len(buf))
		//	println("DEBUG: content=>", string(buf))
		bu = b
		return true
	}

	easy.Setopt(curl.OPT_WRITEFUNCTION, fooTest)
	easy.Setopt(curl.OPT_CUSTOMREQUEST, "{\"method\":\"getstat\"}")
	if err := easy.Perform(); err != nil {
		fmt.Printf("ERROR: %v\n", err)
	}
	e := parseOutput(bu)
	for _, v := range e.Results {
		hrtotal += float64(v.Hashrate)
		numMiners++
	}
	o := output.NewOutput()
	o.Hashrate = hrtotal
	o.NumMiners = numMiners
	js, _ := json.Marshal(o)
	*buf = js
}
