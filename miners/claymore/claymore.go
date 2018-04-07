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

func newClaymoreOutput() *claymoreOutput {
	c := new(claymoreOutput)
	c.Result = make([]string, 9, 20)
	return c
}

func HitClaymore(host_l string, minerPort_l string, buf *[]byte) {
	cmd := "{\"method\":\"miner_getstat1\"}"
	resp, err := dialminer.DialMiner(host_l, minerPort_l, cmd)
	if err != nil {
		*buf = output.MakeJSONError("claymore", err)
		return
	}
	resp = bytes.Trim(resp, "\x00")

	c := newClaymoreOutput()
	err = json.Unmarshal(resp, &c)
	if err != nil {
		*buf = output.MakeJSONError("claymore", err)
		return
	}
	s := strings.Split(c.Result[2], ";")[0]
	hrtotal, err := strconv.Atoi(s)
	if err != nil {
		*buf = output.MakeJSONError("claymore", err)
		return
	}

	numMiners := len(strings.Split(c.Result[3], ";"))

	hrstring := strconv.FormatFloat(float64(hrtotal), 'f', -1, 64) + " MH/s"
	js, err := output.MakeJSON_full("claymore", float64(hrtotal), hrstring, numMiners, 0)
	if err != nil {
		*buf = output.MakeJSONError("claymore", err)
		return
	}
	*buf = js
}
