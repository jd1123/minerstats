package ethminer

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"bitbucket.org/minerstats/dialminer"
	"bitbucket.org/minerstats/output"
)

type ethminerOutput struct {
	Id      int      `json:"id"`
	Result  []string `json:"result"`
	JSONRpc string   `json:"json_rpc"`
}

func newEthminerOutput() *ethminerOutput {
	c := new(ethminerOutput)
	c.Result = make([]string, 9, 20)
	return c
}

func HitEthminer(host_l string, minerPort_l string, buf *[]byte) {
	cmd := "{\"id\":0,\"jsonrpc\":\"2.0\",\"method\":\"miner_getstat1\"}\n"
	resp, err := dialminer.DialMiner(host_l, minerPort_l, cmd)
	if err != nil {
		*buf = output.MakeJSONError("ethminer", err)
		return
	}
	resp = bytes.Trim(resp, "\x00")
	c := newEthminerOutput()
	err = json.Unmarshal(resp, &c)
	if err != nil {
		*buf = output.MakeJSONError("ethminer", err)
		return
	}
	s := strings.Split(c.Result[2], ";")[0]
	hrtotal, err := strconv.Atoi(s)
	if err != nil {
		*buf = output.MakeJSONError("ethminer", err)
		return
	}

	numMiners := len(strings.Split(c.Result[3], ";"))

	hrstring := strconv.FormatFloat(float64(hrtotal), 'f', 2, 64) + " MH/s"

	js, err := output.MakeJSON_full("ethminer", float64(hrtotal), hrstring, numMiners, 0)
	if err != nil {
		*buf = output.MakeJSONError("ethminer", err)
		return
	}
	*buf = js
}
