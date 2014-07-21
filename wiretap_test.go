package wiretap_test

import (
	"fmt"
	"testing"

	"github.com/tmc/wiretap"
)

func TestBasics(t *testing.T) {
	w, err := wiretap.OpenOffline("test.pcap")
	fmt.Println(w, err)
	ph, err := w.PacketHeader()
	fmt.Println(ph, err)
	fmt.Println(ph.Encapsulation())
}
