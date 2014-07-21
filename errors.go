package wiretap

/*
#cgo pkg-config: glib-2.0
#cgo CFLAGS: -I/usr/local/include/wireshark -I/usr/include/wireshark
#cgo LDFLAGS: -lwiretap

#include "wiretap/wtap.h"

*/
import "C"
import (
	"errors"
	"fmt"
)

type Error struct {
	Code int
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message())
}
func (e Error) Message() string {
	return C.GoString(C.wtap_strerror(C.int(e.Code)))
}

func mkErr(code int) error {
	if err, ok := errorCodeToError[code]; ok {
		return err
	}
	return ErrUnknown
}

var (
	ErrNotRegularFile            = Error{Code: C.WTAP_ERR_NOT_REGULAR_FILE}
	ErrRandomOpenPipe            = Error{Code: C.WTAP_ERR_RANDOM_OPEN_PIPE}
	ErrFileUnknownFormat         = Error{Code: C.WTAP_ERR_FILE_UNKNOWN_FORMAT}
	ErrUnsupported               = Error{Code: C.WTAP_ERR_UNSUPPORTED}
	ErrCantWriteToPipe           = Error{Code: C.WTAP_ERR_CANT_WRITE_TO_PIPE}
	ErrCantOpen                  = Error{Code: C.WTAP_ERR_CANT_OPEN}
	ErrUnsupportedFileType       = Error{Code: C.WTAP_ERR_UNSUPPORTED_FILE_TYPE}
	ErrUnsupportedEncap          = Error{Code: C.WTAP_ERR_UNSUPPORTED_ENCAP}
	ErrEncapPerPacketUnsupported = Error{Code: C.WTAP_ERR_ENCAP_PER_PACKET_UNSUPPORTED}
	ErrCantClose                 = Error{Code: C.WTAP_ERR_CANT_CLOSE}
	ErrCantRead                  = Error{Code: C.WTAP_ERR_CANT_READ}
	ErrShortRead                 = Error{Code: C.WTAP_ERR_SHORT_READ}
	ErrBadFile                   = Error{Code: C.WTAP_ERR_BAD_FILE}
	ErrShortWrite                = Error{Code: C.WTAP_ERR_SHORT_WRITE}
	ErrUncTruncated              = Error{Code: C.WTAP_ERR_UNC_TRUNCATED}
	ErrUncOverflow               = Error{Code: C.WTAP_ERR_UNC_OVERFLOW}
	ErrUncBadOffset              = Error{Code: C.WTAP_ERR_UNC_BAD_OFFSET}
	ErrRandomOpenStdin           = Error{Code: C.WTAP_ERR_RANDOM_OPEN_STDIN}
	ErrCompressionNotSupported   = Error{Code: C.WTAP_ERR_COMPRESSION_NOT_SUPPORTED}
	ErrCantSeek                  = Error{Code: C.WTAP_ERR_CANT_SEEK}
	ErrCantSeekCompressed        = Error{Code: C.WTAP_ERR_CANT_SEEK_COMPRESSED}
	ErrDecompress                = Error{Code: C.WTAP_ERR_DECOMPRESS}
	ErrInternal                  = Error{Code: C.WTAP_ERR_INTERNAL}
	ErrUnknown                   = errors.New("unknown errror")
)

var (
	allErrors        []Error
	errorCodeToError map[int]Error
)

func init() {
	allErrors = []Error{
		ErrNotRegularFile,
		ErrRandomOpenPipe,
		ErrFileUnknownFormat,
		ErrUnsupported,
		ErrCantWriteToPipe,
		ErrCantOpen,
		ErrUnsupportedFileType,
		ErrUnsupportedEncap,
		ErrEncapPerPacketUnsupported,
		ErrCantClose,
		ErrCantRead,
		ErrShortRead,
		ErrBadFile,
		ErrShortWrite,
		ErrUncTruncated,
		ErrUncOverflow,
		ErrUncBadOffset,
		ErrRandomOpenStdin,
		ErrCompressionNotSupported,
		ErrCantSeek,
		ErrCantSeekCompressed,
		ErrDecompress,
		ErrInternal,
	}
	errorCodeToError = map[int]Error{}
	for _, err := range allErrors {
		errorCodeToError[err.Code] = err
	}
}
