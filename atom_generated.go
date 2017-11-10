// This file was autogenerated using go run mkcode.go -- atom.go
// DO NOT EDIT.

package ntdll

import "unsafe"

// The AtomInformationClass constants have been derived from the ATOM_INFORMATION_CLASS enum definition.
type AtomInformationClass uint32

const (
	AtomBasicInformation AtomInformationClass = 0
	AtomTableInformation                      = 1
)

var (
	procNtAddAtom              = modntdll.NewProc("NtAddAtom")
	procNtDeleteAtom           = modntdll.NewProc("NtDeleteAtom")
	procNtFindAtom             = modntdll.NewProc("NtFindAtom")
	procNtQueryInformationAtom = modntdll.NewProc("NtQueryInformationAtom")
)

// AtomBasicInformationT has been derived from the ATOM_BASIC_INFORMATION struct definition.
type AtomBasicInformationT struct {
	UsageCount uint16
	Flags      uint16
	NameLength uint16
	Name       [1]uint16
}

// AtomTableInformationT has been derived from the ATOM_TABLE_INFORMATION struct definition.
type AtomTableInformationT struct {
	NumberOfAtoms uint32
	Atoms         [1]RtlAtom
}

// OUT-parameter: Atom.
func NtAddAtom(AtomName *uint16,
	Length uint32,
	Atom *RtlAtom) NtStatus {
	r0, _, _ := procNtAddAtom.Call(uintptr(unsafe.Pointer(AtomName)),
		uintptr(Length),
		uintptr(unsafe.Pointer(Atom)))
	return NtStatus(r0)
}

func NtDeleteAtom(Atom RtlAtom) NtStatus {
	r0, _, _ := procNtDeleteAtom.Call(uintptr(Atom))
	return NtStatus(r0)
}

// OUT-parameter: Atom.
// *OPT-parameter: Atom.
func NtFindAtom(AtomName *uint16,
	Length uint32,
	Atom *RtlAtom) NtStatus {
	r0, _, _ := procNtFindAtom.Call(uintptr(unsafe.Pointer(AtomName)),
		uintptr(Length),
		uintptr(unsafe.Pointer(Atom)))
	return NtStatus(r0)
}

// OUT-parameter: AtomInformation, ReturnLength.
func NtQueryInformationAtom(Atom RtlAtom,
	AtomInformationClass AtomInformationClass,
	AtomInformation *byte,
	AtomInformationLength uint32,
	ReturnLength *uint32) NtStatus {
	r0, _, _ := procNtQueryInformationAtom.Call(uintptr(Atom),
		uintptr(AtomInformationClass),
		uintptr(unsafe.Pointer(AtomInformation)),
		uintptr(AtomInformationLength),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}
