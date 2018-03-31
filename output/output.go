package output

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
