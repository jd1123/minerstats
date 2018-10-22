package trex

import (
	"bytes"
	"strconv"
	"encoding/json"
	"github.com/jd1123/minerstats/dialminer"
	"github.com/jd1123/minerstats/output"
)

type Gpu struct {
  Power float64 `json:"power"`
}

type trexOutput struct {
  Gpus []Gpu `json:"gpus"`
  Hashrate float64 `json:"hashrate"`
}

// Uses JSON response to telnet API, so use -J -b
func HitTrex(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	var totalPower float64 = 0

	resp, err := dialminer.DialMiner(host_l, minerPort_l, "summary\n")
	if err != nil {
		*buf = output.MakeJSONError("trex", err)
		return
	}

	resp = bytes.Trim(resp, "\x00")
  t := new(trexOutput)
  err = json.Unmarshal(resp, &t)
	if err != nil {
		*buf = output.MakeJSONError("trex", err)
		return
	}

  //m := make(map[string]string)
  //subtokens := strings.Split(string(resp), ";")
  //for j := 0; j < len(subtokens)-1; j++ {
  //  sst := strings.Split(subtokens[j], "=")
  //  m[sst[0]] = sst[1]
  //}
  //  hr, _ := strconv.ParseFloat(m["KHS"], 64)
  //  power, _ := strconv.ParseFloat(m["POWER"], 64)
  //  power = power / 1000
  //  hrtotal += hr * 1000
  //  totalPower += power
  //  numMiners++
  //}
  hrtotal = t.Hashrate
  numMiners = len(t.Gpus)
  for i := 0; i < numMiners; i++ {
    totalPower += t.Gpus[i].Power
  }
	hrstring := strconv.FormatFloat(hrtotal/1000000, 'f', 2, 64) + " MH/s"
	js, err := output.MakeJSON_full("trex", hrtotal, hrstring, numMiners, totalPower)
	if err != nil {
		*buf = output.MakeJSONError("trex", err)
		return
	}

	*buf = js
}
