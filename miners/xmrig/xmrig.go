package xmrig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/minerstats/output"
)

var defaultStruct map[string]interface{}

func HitXMRig(host_l string, port_l string, buf *[]byte) {
	fullhost := "http://" + host_l + ":" + port_l
	resp, err := http.Get(fullhost)
	if err != nil {
		fmt.Println("xmrig api error!", err)
		*buf = []byte("connection error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &defaultStruct)
	hrtotal := defaultStruct["hashrate"].(map[string]interface{})["total"].([]interface{})[0].(float64)
	numMiners := len(defaultStruct["health"].([]interface{}))
	hrstring := strconv.FormatFloat(hrtotal, 'f', -1, 64) + " H/s"
	js, err := output.MakeJSON_full("xmrig-nvidia", hrtotal, hrstring, numMiners, 0)
	if err != nil {
		panic(err)
	}
	*buf = js

}
