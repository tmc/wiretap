package wiretap

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/include/wireshark -I/usr/include/wireshark
#cgo LDFLAGS: -lwiretap

#include "wiretap/wtap.h"

*/
import "C"
import "fmt"

type EncapsulationType int

func (et EncapsulationType) String() string {
	return C.GoString(C.wtap_encap_string(C.int(et)))
}

func (et EncapsulationType) ShortString() string {
	return C.GoString(C.wtap_encap_short_string(C.int(et)))
}

func NewEncapsulationTypeFromString(encap string) (EncapsulationType, error) {
	et := int(C.wtap_short_string_to_encap(C.CString(encap)))
	if et == -1 {
		return -1, fmt.Errorf("Invalid encapsulation type '%s'", encap)
	}
	return EncapsulationType(et), nil
}
