package output

import (
	"encoding/json"
)

type Output struct {
	Minername      string  `json:"minername"`
	Hashrate       float64 `json:"hashrate"`
	HashrateString string  `json:"string_hashrate"`
	NumMiners      int     `json:"numminers"`
	TotalPower     float64 `json:"power"`
	Error          string  `json:"error"`
}

type OutputEntry struct {
	ID       string  `json:"id"`
	Hashrate float64 `json:"hashrate"`
}

func NewOutput() *Output {
	o := new(Output)
	return o
}

func ErrorOutput(miner string, err error) *Output {
	o := new(Output)
	o.Minername = miner
	o.Error = err.Error()
	return o
}

func MakeJSON(minerName string, hrtotal float64, numMiners int) ([]byte, error) {
	o := NewOutput()
	o.Minername = minerName
	o.Hashrate = hrtotal
	o.NumMiners = numMiners
	o.TotalPower = -1
	js, err := json.Marshal(o)
	return js, err
}

func MakeJSONError(minerName string, err error) []byte {
	o := ErrorOutput(minerName, err)
	js, err_l := json.Marshal(o)
	if err_l != nil {
		panic(err_l)
	}
	return js
}

func MakeJSON_full(minerName string, hrtotal float64, hrstring string, numMiners int, totalPower float64) ([]byte, error) {
	o := NewOutput()
	o.Minername = minerName
	o.Hashrate = hrtotal
	o.HashrateString = hrstring
	o.NumMiners = numMiners
	o.TotalPower = totalPower
	js, err := json.Marshal(o)
	return js, err
}
