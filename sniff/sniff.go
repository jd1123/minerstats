package sniff

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type SniffOut struct {
}

func SniffMiner() (string, error) {
	fmt.Println("sniffing")
	out, err := exec.Command("netstat", "-plnt").Output()
	if err != nil {
		return "", err
	}

	out_t := strings.Split(string(out), "\n")
	removeSpaces := regexp.MustCompile(`[\s\p{Zs}]{2,}`)

	for i := 2; i < len(out_t)-1; i++ {
		fmt.Println(removeSpaces.ReplaceAllString(out_t[i], ";"))
	}

	return "", nil
}
