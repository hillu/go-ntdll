// This file was autogenerated using go run mkcode.go -- security.go
// DO NOT EDIT.

package ntdll

import "unsafe"
import "reflect"

var (
	procNtQuerySecurityObject = modntdll.NewProc("NtQuerySecurityObject")
	procNtSetSecurityObject   = modntdll.NewProc("NtSetSecurityObject")
)

// Acl has been derived from the ACL struct definition.
type Acl struct {
	AclRevision byte
	Sbz1        byte
	AclSize     uint16
	AceCount    uint16
	Sbz2        uint16
}

// Sid has been derived from the SID struct definition.
type Sid struct {
	Revision            byte
	SubAuthorityCount   byte
	IdentifierAuthority SidIdentifierAuthority
	SubAuthority        [1]uint32
}

// SubAuthoritySlice returns a slice over the elements of Sid.SubAuthority
func (t *Sid) SubAuthoritySlice(size int) []uint32 {
	s := []uint32{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&t.SubAuthority[0]))
	hdr.Len = size
	hdr.Cap = size
	return s
}

// SidIdentifierAuthority has been derived from the SID_IDENTIFIER_AUTHORITY struct definition.
type SidIdentifierAuthority struct {
	Value [6]byte
}

// SecurityDescriptor has been derived from the SECURITY_DESCRIPTOR struct definition.
type SecurityDescriptor struct {
	Revision byte
	Sbz1     byte
	Control  SecurityDescriptorControl
	Owner    *Sid
	Group    *Sid
	Sacl     *Acl
	Dacl     *Acl
}

// unknown-parameter: Handle, SecurityInformation, SecurityDescriptor, Length, LengthNeeded.
func NtQuerySecurityObject(
	Handle Handle,
	SecurityInformation SecurityInformationT,
	SecurityDescriptor *SecurityDescriptor,
	Length uint32,
	LengthNeeded *uint32,
) NtStatus {
	r0, _, _ := procNtQuerySecurityObject.Call(uintptr(Handle),
		uintptr(SecurityInformation),
		uintptr(unsafe.Pointer(SecurityDescriptor)),
		uintptr(Length),
		uintptr(unsafe.Pointer(LengthNeeded)))
	return NtStatus(r0)
}

// unknown-parameter: Handle, SecurityInformation, SecurityDescriptor.
func NtSetSecurityObject(
	Handle Handle,
	SecurityInformation SecurityInformationT,
	SecurityDescriptor *SecurityDescriptor,
) NtStatus {
	r0, _, _ := procNtSetSecurityObject.Call(uintptr(Handle),
		uintptr(SecurityInformation),
		uintptr(unsafe.Pointer(SecurityDescriptor)))
	return NtStatus(r0)
}
