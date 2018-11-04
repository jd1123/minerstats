package xmrig

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jd1123/minerstats/output"
)

var defaultStruct map[string]interface{}

func HitXMRig(host_l string, port_l string, buf *[]byte) {
	fullhost := "http://" + host_l + ":" + port_l
	resp, err := http.Get(fullhost)
	if err != nil {
		*buf = output.MakeJSONError("xmrig", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*buf = output.MakeJSONError("xmrig", err)
		return
	}
	json.Unmarshal(body, &defaultStruct)
  hrs := defaultStruct["hashrate"].(map[string]interface{})
	hrtotal := hrs["total"].([]interface{})[0].(float64)
	numMiners := len(hrs)
	hrstring := strconv.FormatFloat(hrtotal, 'f', 2, 64) + " H/s"
	js, err := output.MakeJSON_full("xmrig", hrtotal, hrstring, numMiners, 0)
	if err != nil {
		*buf = output.MakeJSONError("xmrig", err)
		return
	}
	*buf = js

}
