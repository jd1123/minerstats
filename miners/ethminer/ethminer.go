package ethminer

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type claymoreResult struct {
	id0 string
	id1 string
	id2 string
	id3 string
	id4 string
	id5 string
	id6 string
	id7 string
	id8 string
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
	fmt.Println(string(resp))
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
