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
// Example:
/*
buf := make([]byte, 128)
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
	for {
		if st := fn(); st == STATUS_BUFFER_OVERFLOW || st == STATUS_BUFFER_TOO_SMALL {
			if int(*resultLength) <= cap(*buf) {
				(*reflect.SliceHeader)(unsafe.Pointer(buf)).Len = int(*resultLength)
			} else {
				*buf = make([]byte, int(*resultLength))
			}
			continue
		} else {
			return st
		}
	}
}
