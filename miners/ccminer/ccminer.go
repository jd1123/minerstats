package ccminer

import (
	"strconv"
	"strings"

	"bitbucket.org/minerstats/dialminer"
	"bitbucket.org/minerstats/output"
)

func HitCCMiner(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	var totalPower float64 = 0

	resp, err := dialminer.DialMiner(host_l, minerPort_l, "threads")
	if err != nil {
		*buf = output.MakeJSONError("ccminer", err)
	}

	threadsinfo := string(resp)
	tokens := strings.Split(threadsinfo, "|")
	for i := 0; i < len(tokens)-1; i++ {
		m := make(map[string]string)
		subtokens := strings.Split(tokens[i], ";")
		for j := 0; j < len(subtokens)-1; j++ {
			sst := strings.Split(subtokens[j], "=")
			m[sst[0]] = sst[1]
		}
		ix := m["GPU"]
		if ix == "" {
		} else {
			hr, _ := strconv.ParseFloat(m["KHS"], 64)
			power, _ := strconv.ParseFloat(m["POWER"], 64)
			power = power / 1000
			hrtotal += hr * 1000
			totalPower += power
			numMiners++
		}
	}
	hrstring := strconv.FormatFloat(hrtotal/1000000, 'f', 2, 64) + " MH/s"
	js, err := output.MakeJSON_full("ccminer", hrtotal, hrstring, numMiners, totalPower)
	if err != nil {
		*buf = output.MakeJSONError("ccminer", err)
		return
	}

	*buf = js
}
