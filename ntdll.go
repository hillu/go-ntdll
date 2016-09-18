package ntdll

import (
	"golang.org/x/sys/windows"
)

var modntdll = windows.NewLazyDLL("ntdll")

type Handle windows.Handle

func fromBool(b bool) (r uintptr) {
	if b {
		r = 1
	}
	return
}
