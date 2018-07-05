package zm

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/jd1123/minerstats/dialminer"
	"github.com/jd1123/minerstats/output"
)

type result struct {
	Gpuid          int
	Cudaid         int
	Busid          string
	Name           string
	GpuStatus      int
	Solver         int
	Temp           int
	GPUPower       float64 `json:"power_usage"`
	Hashrate       float64 `json:"avg_sol_ps"`
	AcceptedShares int
	RejectedShares int
	StartTime      int
}

type zmOutput struct {
	Id        int      `json:"id"`
	Uptime    int      `json:"uptime"`
	Contime   int      `json:"contime"`
	Server    string   `json:"server"`
	Port      int      `json:"port"`
	User      string   `json:"user"`
	Version   string   `json:"version"`
	Error     string   `json:"error"`
	StartTime int      `json:"start_time"`
	Results   []result `json:"result"`
}

func newZMOutput() *zmOutput {
	z := new(zmOutput)
	z.Results = make([]result, 15, 30)
	return z
}

func HitZM(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	var totalPower float64 = 0

	resp, err := dialminer.DialMiner(host_l, minerPort_l, "{\"method\":\"getstat\"}")
	if err != nil {
		panic(err)
	}
	resp = bytes.Trim(resp, "\x00")

	z := newZMOutput()
	err = json.Unmarshal(resp, &z)

	// ZM for some reason sends malformed JSON on
	// certain requests. I am not sure if this is a local only
	// problem (as I have network issues at the moment) or something
	// deeper.
	if err != nil {
		*buf = output.MakeJSONError("zm", err)
		return
	}

	for _, v := range z.Results {
		numMiners++
		hrtotal += float64(v.Hashrate)
		totalPower += float64(v.GPUPower)
	}

	hrstring := strconv.FormatFloat(hrtotal, 'f', 2, 64) + " Sol/s"

	o := output.NewOutput()
	o.Minername = "zm"
	o.Hashrate = hrtotal
	o.HashrateString = hrstring
	o.NumMiners = numMiners
	o.TotalPower = totalPower

	js, err := json.Marshal(o)
	if err != nil {
		*buf = output.MakeJSONError("zm", err)
		return
	}

	*buf = js
}
