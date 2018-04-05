package sniff

import "testing"

func TestValidMiner(t *testing.T) {
	v := newValidMiner("ccminer", "4068")
	v.PrintMiner()
}
