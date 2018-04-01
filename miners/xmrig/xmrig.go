package xmrig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	js, err := output.MakeJSON("xmrig-nvidia", hrtotal, numMiners)
	if err != nil {
		panic(err)
	}
	*buf = js

}
