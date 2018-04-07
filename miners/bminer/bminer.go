package bminer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"bitbucket.org/minerstats/output"
)

type BminerUtilization struct {
	GPU    int `json:"gpu"`
	Memory int `json:"memory"`
}

type BminerClocks struct {
	Core   int `json:"core"`
	Memory int `json:"memory"`
}

type BminerPCI struct {
	Barl int `json:"barl_used"`
	RX   int `json:"rx_throughput"`
	TX   int `json:"tx_throughput"`
}

type BminerSolver struct {
	SolutionRate float64 `json:"solution_rate"`
	NonceRate    float64 `json:"nonce_rate"`
}

type BminerDevice struct {
	Temp             int               `json:"temperature"`
	Power            int               `json:"power"`
	GlobalMemoryUsed int               `json:"global_memory_used"`
	Utilization      BminerUtilization `json:"utilization"`
	Clocks           BminerClocks      `json:"clocks"`
	PCI              BminerPCI         `json:"pci"`
}

type BminerMiner struct {
	Solver BminerSolver `json:"solver"`
	Device BminerDevice `json:"device"`
}

type BminerStratum struct {
	AcceptedShares    int     `json:"accepted_shares"`
	RejectedShares    int     `json:"rejected_shares"`
	AcceptedShareRate float64 `json:"accepted_share_rate"`
	RejectedShareRate float64 `json:"rejected_share_rate"`
}

type BminerJSON struct {
	Stratum   BminerStratum          `json:"stratum"`
	Miners    map[string]BminerMiner `json:"miners"`
	Version   string                 `json:"version"`
	StartTime int                    `json:"start_time"`
}

func newBminerJSON() *BminerJSON {
	bj := new(BminerJSON)
	bj.Miners = make(map[string]BminerMiner)
	return bj
}

func strToInt(s string) int {
	kint, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("i=%d, type: %T\n", kint, kint)
		panic(err)
	}
	return kint
}

func parseBminerOutput(b []byte) *BminerJSON {
	var bminerJson = newBminerJSON()
	json.Unmarshal(b, &bminerJson)
	return bminerJson
}

/*
func HitBminer(host_l string, minerPort_l string, buf *[]byte) {
	fullhost := "http://" + host_l + ":" + minerPort_l + "/api/status"
	resp, err := http.Get(fullhost)
	if err != nil {
		fmt.Println("bminer api error!", err)
		*buf = []byte("connection error")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body read error!", err)
	}

	j := parseBminerOutput(body)
	if err != nil {
		panic(err)
	}
	js, _ := json.Marshal(j)
	*buf = js
}
*/
func HitBminer(host_l string, minerPort_l string, buf *[]byte) {
	var hrtotal float64 = 0
	var numMiners int = 0
	fullhost := "http://" + host_l + ":" + minerPort_l + "/api/status"
	resp, err := http.Get(fullhost)
	if err != nil {
		//*buf = []byte("connection error")
		*buf = output.MakeJSONError("bminer", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		*buf = output.MakeJSONError("bminer", err)
		return
	}

	j := parseBminerOutput(body)
	for _, v := range j.Miners {
		numMiners++
		hrtotal += v.Solver.SolutionRate
	}
	o := output.NewOutput()
	o.Minername = "bminer"
	o.Hashrate = hrtotal
	o.HashrateString = strconv.FormatFloat(hrtotal, 'f', 2, 64) + " Sols/s"
	o.NumMiners = numMiners

	js, err := json.Marshal(o)
	if err != nil {
		*buf = output.MakeJSONError("bminer", err)
		return
	}
	*buf = js
}
