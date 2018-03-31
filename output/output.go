package output

import (
	"encoding/json"
)

type Output struct {
	Minername string  `json:"minername"`
	Hashrate  float64 `json:"hashrate"`
	NumMiners int     `json:"numminers"`
}

type OutputEntry struct {
	ID       string  `json:"id"`
	Hashrate float64 `json:"hashrate"`
}

func NewOutput() *Output {
	o := new(Output)
	return o
}

func MakeJSON(minerName string, hrtotal float64, numMiners int) ([]byte, error) {
	o := NewOutput()
	o.Minername = "ewbf"
	o.Hashrate = hrtotal
	o.NumMiners = numMiners
	js, err := json.Marshal(o)
	return js, err
}
