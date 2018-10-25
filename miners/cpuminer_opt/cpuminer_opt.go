package cpuminer_opt

import (
	"strconv"
	"strings"

	"github.com/jd1123/minerstats/dialminer"
	"github.com/jd1123/minerstats/output"
)

func HitCPUMinerOpt(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	var totalPower float64 = 0

	resp, err := dialminer.DialMiner(host_l, minerPort_l, "threads")
	if err != nil {
		*buf = output.MakeJSONError("cpuminer_opt", err)
	}

	threadsinfo := string(resp)
	cores := strings.Split(threadsinfo, "|")
	for i := 0; i < len(cores)-1; i++ {
    attrs := strings.Split(cores[i], ";")
    hrAttr := strings.Split(attrs[1], "=")
    hr, _ := strconv.ParseFloat(hrAttr[1], 64)
    hrtotal += hr
    numMiners++
	}
	hrstring := strconv.FormatFloat(hrtotal/1000, 'f', 2, 64) + " KH/s"
	js, err := output.MakeJSON_full("cpuminer_opt", hrtotal, hrstring, numMiners, totalPower)
	if err != nil {
		*buf = output.MakeJSONError("cpuminer_opt", err)
		return
	}

	*buf = js
}
