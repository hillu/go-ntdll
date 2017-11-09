package ntdll

//go:generate -command statuscrawler go run statuscrawler.go --
//go:generate statuscrawler

import (
	"fmt"
)

type NtStatus uint32

func (s NtStatus) Error() error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("NtStatus 0x%08x (%s)", uint32(s), s)
}

func (s NtStatus) IsSuccess() bool {
	return s>>30 == 0
}

func (s NtStatus) IsInofmational() bool {
	return s>>30 == 1
}

func (s NtStatus) IsWarning() bool {
	return s>>30 == 2
}

func (s NtStatus) IsError() bool {
	return s>>30 == 3
}

func (s NtStatus) Facility() uint32 {
	return (uint32(s) >> 16) & 0x0fff
}

func (s NtStatus) Code() uint32 {
	return uint32(s) & 0xffff
}
