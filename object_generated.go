// This file was autogenerated using go run mkcode.go -- object.go
// DO NOT EDIT.

package ntdll

import "unsafe"

var (
	procNtOpenDirectoryObject      = modntdll.NewProc("NtOpenDirectoryObject")
	procNtQueryDirectoryObject     = modntdll.NewProc("NtQueryDirectoryObject")
	procNtOpenSymbolicLinkObject   = modntdll.NewProc("NtOpenSymbolicLinkObject")
	procNtQuerySymbolicLinkObject  = modntdll.NewProc("NtQuerySymbolicLinkObject")
	procNtCreateSymbolicLinkObject = modntdll.NewProc("NtCreateSymbolicLinkObject")
	procNtCreateDirectoryObject    = modntdll.NewProc("NtCreateDirectoryObject")
)

// ObjectAttributes has been derived from the OBJECT_ATTRIBUTES struct definition.
type ObjectAttributes struct {
	Length                   uint32
	RootDirectory            Handle
	ObjectName               *UnicodeString
	Attributes               uint32
	SecurityDescriptor       *byte
	SecurityQualityOfService *byte
}

// ObjectDirectoryInformationT has been derived from the OBJECT_DIRECTORY_INFORMATION struct definition.
type ObjectDirectoryInformationT struct {
	Name     UnicodeString
	TypeName UnicodeString
}

// OUT-parameter: DirectoryHandle.
func NtOpenDirectoryObject(DirectoryHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes) NtStatus {
	r0, _, _ := procNtOpenDirectoryObject.Call(uintptr(unsafe.Pointer(DirectoryHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

// OUT-parameter: Buffer, ReturnLength.
// INOUT-parameter: Context.
// *OPT-parameter: Buffer, ReturnLength.
func NtQueryDirectoryObject(DirectoryHandle Handle,
	Buffer *byte,
	Length uint32,
	ReturnSingleEntry bool,
	RestartScan bool,
	Context *uint32,
	ReturnLength *uint32) NtStatus {
	r0, _, _ := procNtQueryDirectoryObject.Call(uintptr(DirectoryHandle),
		uintptr(unsafe.Pointer(Buffer)),
		uintptr(Length),
		fromBool(ReturnSingleEntry),
		fromBool(RestartScan),
		uintptr(unsafe.Pointer(Context)),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}

// OUT-parameter: LinkHandle.
func NtOpenSymbolicLinkObject(LinkHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes) NtStatus {
	r0, _, _ := procNtOpenSymbolicLinkObject.Call(uintptr(unsafe.Pointer(LinkHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

// OUT-parameter: ReturnedLength.
// INOUT-parameter: LinkTarget.
// *OPT-parameter: ReturnedLength.
func NtQuerySymbolicLinkObject(LinkHandle Handle,
	LinkTarget *UnicodeString,
	ReturnedLength *uint32) NtStatus {
	r0, _, _ := procNtQuerySymbolicLinkObject.Call(uintptr(LinkHandle),
		uintptr(unsafe.Pointer(LinkTarget)),
		uintptr(unsafe.Pointer(ReturnedLength)))
	return NtStatus(r0)
}

// OUT-parameter: SymbolicLinkHandle.
func NtCreateSymbolicLinkObject(SymbolicLinkHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	TargetName *UnicodeString) NtStatus {
	r0, _, _ := procNtCreateSymbolicLinkObject.Call(uintptr(unsafe.Pointer(SymbolicLinkHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(TargetName)))
	return NtStatus(r0)
}

// OUT-parameter: DirectoryHandle.
func NtCreateDirectoryObject(DirectoryHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes) NtStatus {
	r0, _, _ := procNtCreateDirectoryObject.Call(uintptr(unsafe.Pointer(DirectoryHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}
