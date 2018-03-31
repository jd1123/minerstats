package zm

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
	Hashrate       float64 `json:"avg_sol_ps"`
	AcceptedShares int
	RejectedShares int
	StartTime      int
}

type ZMOutput struct {
	Id        int      `json:"id"`
	Uptime    int      `json:"uptime"`
	Contime   int      `json:"contime"`
	Server    string   `json:"server"`
	Port      int      `json:"port"`
	User      string   `json:"user"`
	Version   string   `json:"version"`
	Error     string   `json:"error"`
	StartTime int      `json:"start_time"`
	Results   []Result `json:"result"`
}

func NewZMOutput() *ZMOutput {
	z := new(ZMOutput)
	z.Results = make([]Result, 15, 30)
	return z
}

func HitZM(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0

	resp, err := dialminer.DialMiner(host_l, minerPort_l, "{\"method\":\"getstat\"}")
	if err != nil {
		panic(err)
	}
	resp = bytes.Trim(resp, "\x00")

	z := NewZMOutput()
	err = json.Unmarshal(resp, &z)

	// This is made to intentionally break the program
	// on this error. ZM for some reason sends malformed JSON on
	// certain requests. I am not sure if this is a local only
	// problem (as I have network issues at the moment) or something
	// deeper.
	if err != nil {
		*buf = []byte("error:" + err.Error())
		return
	}

	for _, v := range z.Results {
		numMiners++
		hrtotal += float64(v.Hashrate)
	}
	o := output.NewOutput()
	o.Minername = "ZM"
	o.Hashrate = hrtotal
	o.NumMiners = numMiners
	js, _ := json.Marshal(o)
	*buf = js
}
