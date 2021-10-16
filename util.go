package ntdll

import (
	"reflect"
	"unsafe"
)

// CallWithExpandingBuffer calls fn which encapsulates the actual API
// call using buf and resultLength. If fn returns a status that
// indicates a too small buffer, the buffer is expanded and the call
// retried, until it succeeds (or fails for another reason).
//
// If the API call succeeds, the buffer size is adjusted according to
// resultLength.
//
// Example:
/*
var buf []byte
var rlen uint32
if st := CallWithExpandingBuffer(func() NtStatus {
    return NtQueryKey(
		h, KeyFullInformation,
		&buf[0],
		uint32(len(buf)),
		&rlen,
	)
}, &buf, &rlen); st.IsError() {
	...
}
*/
func CallWithExpandingBuffer(fn func() NtStatus, buf *[]byte, resultLength *uint32) NtStatus {
	if len(*buf) == 0 {
		*buf = make([]byte, 32, 32)
	}
	for {
		st := fn()
		if int(*resultLength) <= cap(*buf) {
			(*reflect.SliceHeader)(unsafe.Pointer(buf)).Len = int(*resultLength)
		} else {
			orig := buf
			*buf = make([]byte, int(*resultLength))
			copy(*buf, *orig)
		}
		if st == STATUS_BUFFER_OVERFLOW || st == STATUS_BUFFER_TOO_SMALL || st == STATUS_INFO_LENGTH_MISMATCH {
			continue
		} else {
			return st
		}
	}
}
