package claymore

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

	"bitbucket.org/minerstats/dialminer"
	"bitbucket.org/minerstats/output"
)

type claymoreOutput struct {
	Id     int      `json:"id"`
	Result []string `json:"result"`
	Error  string   `json:"error"`
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

func NewClaymoreOutput() *claymoreOutput {
	c := new(claymoreOutput)
	c.Result = make([]string, 9, 20)
	return c
}

func HitClaymore(host_l string, minerPort_l string, buf *[]byte) {
	cmd := "{\"method\":\"miner_getstat1\"}"
	resp, err := dialminer.DialMiner(host_l, minerPort_l, cmd)
	if err != nil {
		panic(err)
	}
	resp = bytes.Trim(resp, "\x00")

	c := NewClaymoreOutput()
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

	js, err := output.MakeJSON("claymore", float64(hrtotal), numMiners)
	if err != nil {
		panic(err)
	}
	*buf = js
}
