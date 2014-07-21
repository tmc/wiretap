package wiretap

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/include/wireshark -I/usr/include/wireshark
#cgo LDFLAGS: -lwiretap

#include "wiretap/wtap.h"

*/
import "C"
import "fmt"

type Handle struct {
	wth *C.wtap
}

func (h Handle) SnapLen() int {
	return int(C.wtap_snapshot_length(h.wth))
}

func (h Handle) PacketHeader() (*PacketHeader, error) {
	p := C.wtap_phdr(h.wth)
	if p == nil {
		return nil, fmt.Errorf("invalid packet header")
	}
	return &PacketHeader{p}, nil
}

func (h Handle) FileSize() int {
	return 0
}

func (h Handle) FileType() int {
	return 0
}

func (h Handle) IsCompressed() bool {
	return false
}

func (h Handle) Close() {
}

func (h Handle) FDClose() {
}

func (h Handle) SequentialClose() {
}

func OpenOffline(fileName string) (*Handle, error) {
	h := new(Handle)
	var (
		err     C.int
		errInfo string
	)
	cErrInfo := (*C.gchar)(C.CString(errInfo))

	h.wth = C.wtap_open_offline(C.CString(fileName), &err, &cErrInfo, 1)

	if err != 0 {
		return nil, mkErr(int(err))
	}

	return h, nil
}
