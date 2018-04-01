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
		panic(err)
	}
	resp = bytes.Trim(resp, "\x00")
	c := newEthminerOutput()
	err = json.Unmarshal(resp, &c)
	if err != nil {
		panic(err)
	}
	s := strings.Split(c.Result[2], ";")[0]
	hrtotal, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	numMiners := len(strings.Split(c.Result[3], ";"))

	js, err := output.MakeJSON("ethminer", float64(hrtotal), numMiners)
	if err != nil {
		panic(err)
	}
	*buf = js
}
