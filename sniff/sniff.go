// This package implements the sniffing.
//
// It runs netstat as a subprocess and filters out for known miners.
// A consequence of this is that binaries must be named properly, as we know them
// otherwise it will not detect the miner. The list of known miners will be documented
// and binary names will be specified.

package sniff

import (
	"errors"
	"fmt"
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

type ValidMiner struct {
	Name string
	Port string
}

func (v ValidMiner) PrintMiner() {
	fmt.Println(v.Name + ":" + v.Port)
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

func newValidMiner(name string, port string) *ValidMiner {
	v := new(ValidMiner)
	v.Name = name
	v.Port = port
	return v
}

// This function runs netstat and parses the output.
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

func SniffMiner() ([]*ValidMiner, error) {
	var detectedMiner bool = false
	var foundMiners []*ValidMiner

	procs, err := getNetstatOutput()
	if err != nil {
		return nil, err
	}

	for _, v := range procs {
		if isAMiner(v.name) {
			detectedMiner = true
			foundMiners = append(foundMiners, newValidMiner(v.name, v.port))
		}
	}
	if detectedMiner {
		return foundMiners, nil
	} else {
		return nil, errors.New("No miners found")
	}
}
