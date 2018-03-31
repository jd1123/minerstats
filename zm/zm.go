package zm

import (
	"fmt"

	"bitbucket.org/minerstats/dialminer"
)

func HitZM(host_l string, minerPort_l string, buf *[]byte) {
	resp, err := dialminer.DialMiner(host_l, minerPort_l, "{\"method\":\"getstat\"}")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	//*buf = getHashrate()
}
