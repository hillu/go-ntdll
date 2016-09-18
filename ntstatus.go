package ntdll

import (
	"fmt"
)

type NtStatus uint32

func (s NtStatus) Error() error {
	if s == 0 {
		return nil
	}
	return fmt.Errorf("NtStatus %08x", s)
}
