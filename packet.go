package wiretap

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/include/wireshark -I/usr/include/wireshark
#cgo LDFLAGS: -lwiretap

#include "wiretap/wtap.h"

*/
import "C"

type PacketHeader struct {
	p *C.struct_wtap_pkthdr
}

func (p *PacketHeader) Encapsulation() EncapsulationType {
	return EncapsulationType(int(p.p.pkt_encap))
}
