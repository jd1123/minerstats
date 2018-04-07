// These are old functions and declarations that arent used
// but might be useful in the future. Saved for reference

package ccminer

import (
	"encoding/json"
	"strconv"
	"strings"

	"bitbucket.org/minerstats/dialminer"
)

type HWinfo struct {
	Miners map[string]*HWinfoEntry `json:"miners"`
}

type HWinfoEntry struct {
	Gpu      string  `json:"gpu"`
	Bus      int64   `json:"bus"`
	Card     string  `json:"card"`
	Memclock int64   `json:"mem_clock"`
	Temp     float64 `json:"temp"`
	Power    float64 `json:"power"`
	Fan      int64   `json:"fan"`
}

type Threads struct {
	Miners map[string]*ThreadEntry `json:"miners"`
}

type ThreadEntry struct {
	Gpu   string  `json:"gpu"`
	Bus   int64   `json:"bus"`
	Card  string  `json:"card"`
	Temp  float64 `json:"temp"`
	Fan   int64   `json:"fan"`
	Power float64 `json:"power"`
	KHs   float64 `json:"hashrate"`
	Hs    float64 `json:"rawhashrate"`
	I     float64 `json:"intensity"`
}

func NewThreads() *Threads {
	t := new(Threads)
	t.Miners = make(map[string]*ThreadEntry)
	return t
}
func NewHWinfo() *HWinfo {
	h := new(HWinfo)
	h.Miners = make(map[string]*HWinfoEntry)
	return h
}

func getHwinfo(host string, port string) []byte {
	h := NewHWinfo()
	bu, err := dialminer.DialMiner(host, port, "hwinfo")
	if err != nil {
		return []byte("connection error")
	}
	hwinfo := string(bu)
	tokens := strings.Split(hwinfo, "|")
	for i := 0; i < len(tokens)-1; i++ {
		m := make(map[string]string)
		subtokens := strings.Split(tokens[i], ";")
		for j := 0; j < len(subtokens)-1; j++ {
			sst := strings.Split(subtokens[j], "=")
			m[sst[0]] = sst[1]
		}
		ix := m["GPU"]
		if ix == "" {
		} else {
			h.Miners[ix] = new(HWinfoEntry)
			h.Miners[ix].Gpu = ix
			temp, _ := strconv.ParseFloat(m["TEMP"], 64)
			power, _ := strconv.ParseFloat(m["POWER"], 64)
			fan, _ := strconv.ParseInt(m["FAN"], 10, 32)
			mem, _ := strconv.ParseInt(m["MEM"], 10, 32)
			bus, _ := strconv.ParseInt(m["BUS"], 10, 32)
			h.Miners[ix].Card = m["CARD"]
			h.Miners[ix].Temp = temp
			h.Miners[ix].Memclock = mem
			h.Miners[ix].Fan = fan
			h.Miners[ix].Power = power / 1000
			h.Miners[ix].Bus = bus
		}
	}
	b, _ := json.Marshal(h)
	return b
}

func getThreadsInfo(host string, port string) []byte {
	t := NewThreads()
	bu, err := dialminer.DialMiner(host, port, "threads")
	if err != nil {
		return []byte("connection error")
	}
	threadsinfo := string(bu)
	tokens := strings.Split(threadsinfo, "|")
	for i := 0; i < len(tokens)-1; i++ {
		m := make(map[string]string)
		subtokens := strings.Split(tokens[i], ";")
		for j := 0; j < len(subtokens)-1; j++ {
			sst := strings.Split(subtokens[j], "=")
			m[sst[0]] = sst[1]
		}
		ix := m["GPU"]
		if ix == "" {
		} else {
			t.Miners[ix] = new(ThreadEntry)
			t.Miners[ix].Gpu = ix
			temp, _ := strconv.ParseFloat(m["TEMP"], 64)
			power, _ := strconv.ParseFloat(m["POWER"], 64)
			fan, _ := strconv.ParseInt(m["FAN"], 10, 32)
			bus, _ := strconv.ParseInt(m["BUS"], 10, 32)
			hr, _ := strconv.ParseFloat(m["KHS"], 64)
			intensity, _ := strconv.ParseFloat(m["I"], 64)
			t.Miners[ix].Card = m["CARD"]
			t.Miners[ix].Temp = temp
			t.Miners[ix].Fan = fan
			t.Miners[ix].I = intensity
			t.Miners[ix].KHs = hr
			t.Miners[ix].Hs = hr * 1000
			t.Miners[ix].Power = power / 1000
			t.Miners[ix].Bus = bus
		}
	}
	b, _ := json.Marshal(t)
	return b
}
