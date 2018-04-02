package sniff

import (
	"os/exec"
	"regexp"
	"strings"
)

var miners map[string]string = map[string]string{
	"ccminer": "ccminer",
}

var minerlist []string = []string{
	"ccminer",
	"zm",
	"ewbf",
	"bminer",
	"ethminer",
	"claymore",
	"xmrig",
}

type SniffOut struct {
	miner string
	port  string
}

type netstatProcess struct {
	status string
	port   string
	name   string
}

func isAMiner(a string) bool {
	for _, v := range minerlist {
		if a == v {
			return true
		}
	}
	return false
}

func getNetstatOutput() ([]*netstatProcess, error) {
	out, err := exec.Command("netstat", "-plnt").Output()
	if err != nil {
		return nil, err
	}

	con_to_chk := strings.Split(string(out), "\n")
	removeSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	processes := make([]*netstatProcess, len(con_to_chk)-3)

	for i := 2; i < len(con_to_chk)-1; i++ {
		pidx := i - 2
		processes[pidx] = new(netstatProcess)

		tokens := removeSpaces.ReplaceAllString(con_to_chk[i], ";")

		status := strings.Split(tokens, ";")[4]
		process := strings.Split(tokens, ";")[5]
		port := strings.Split(strings.Split(strings.Split(tokens, ";")[2], " ")[1], ":")[1]

		if process == "-" {
			processes[pidx].name = process
		} else {
			processes[pidx].name = strings.Split(process, "/")[1]
		}

		processes[pidx].status = status
		processes[pidx].port = port
	}
	return processes, nil
}

func SniffMiner() (string, error) {
	var detectedMiner bool = false

	procs, err := getNetstatOutput()
	if err != nil {
		return "", err
	}

	for _, v := range procs {
		if isAMiner(v.name) {
			detectedMiner = true
		}
	}
	if detectedMiner {
		return "yes!", nil
	} else {
		return "no", nil
	}
}
