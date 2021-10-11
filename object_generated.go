// This file was autogenerated using go run mkcode.go -- object.go
// DO NOT EDIT.

package ntdll

import "unsafe"

// The ObjectInformationClass constants have been derived from the OBJECT_INFORMATION_CLASS enum definition.
type ObjectInformationClass uint32

const (
	ObjectBasicInformation ObjectInformationClass = 0
	ObjectNameInformation                         = 1
	ObjectTypeInformation                         = 2
	ObjectAllInformation                          = 3
	ObjectDataInformation                         = 4
)

var (
	procNtOpenDirectoryObject      = modntdll.NewProc("NtOpenDirectoryObject")
	procNtQueryDirectoryObject     = modntdll.NewProc("NtQueryDirectoryObject")
	procNtOpenSymbolicLinkObject   = modntdll.NewProc("NtOpenSymbolicLinkObject")
	procNtQuerySymbolicLinkObject  = modntdll.NewProc("NtQuerySymbolicLinkObject")
	procNtCreateSymbolicLinkObject = modntdll.NewProc("NtCreateSymbolicLinkObject")
	procNtCreateDirectoryObject    = modntdll.NewProc("NtCreateDirectoryObject")
	procNtQueryObject              = modntdll.NewProc("NtQueryObject")
	procNtDuplicateObject          = modntdll.NewProc("NtDuplicateObject")
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

// ObjectBasicInformationT has been derived from the OBJECT_BASIC_INFORMATION struct definition.
type ObjectBasicInformationT struct {
	Attributes               uint32
	GrantedAccess            AccessMask
	HandleCount              uint32
	PointerCount             uint32
	PagedPoolUsage           uint32
	NonPagedPoolUsage        uint32
	Reserved                 [3]uint32
	NameInformationLength    uint32
	TypeInformationLength    uint32
	SecurityDescriptorLength uint32
	CreateTime               int64
}

// ObjectNameInformationT has been derived from the OBJECT_NAME_INFORMATION struct definition.
type ObjectNameInformationT struct {
	Name UnicodeString
}

// GenericMapping has been derived from the GENERIC_MAPPING struct definition.
type GenericMapping struct {
	GenericRead    AccessMask
	GenericWrite   AccessMask
	GenericExecute AccessMask
	GenericAll     AccessMask
}

// ObjectTypeInformationT has been derived from the OBJECT_TYPE_INFORMATION struct definition.
type ObjectTypeInformationT struct {
	TypeName                   UnicodeString
	TotalNumberOfObjects       uint32
	TotalNumberOfHandles       uint32
	TotalPagedPoolUsage        uint32
	TotalNonPagedPoolUsage     uint32
	TotalNamePoolUsage         uint32
	TotalHandleTableUsage      uint32
	HighWaterNumberOfObjects   uint32
	HighWaterNumberOfHandles   uint32
	HighWaterPagedPoolUsage    uint32
	HighWaterNonPagedPoolUsage uint32
	HighWaterNamePoolUsage     uint32
	HighWaterHandleTableUsage  uint32
	InvalidAttributes          uint32
	GenericMapping             GenericMapping
	ValidAccessMask            uint32
	SecurityRequired           bool
	MaintainHandleCount        bool
	TypeIndex                  byte
	ReservedByte               byte
	PoolType                   uint32
	DefaultPagedPoolCharge     uint32
	DefaultNonPagedPoolCharge  uint32
}

// ObjectAllInformationT has been derived from the OBJECT_ALL_INFORMATION struct definition.
type ObjectAllInformationT struct {
	NumberOfObjects       uint32
	ObjectTypeInformation [1]ObjectTypeInformationT
}

// ObjectDataInformationT has been derived from the OBJECT_DATA_INFORMATION struct definition.
type ObjectDataInformationT struct {
	InheritHandle    bool
	ProtectFromClose bool
}

// OUT-parameter: DirectoryHandle.
func NtOpenDirectoryObject(
	DirectoryHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
) NtStatus {
	r0, _, _ := procNtOpenDirectoryObject.Call(uintptr(unsafe.Pointer(DirectoryHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

// OUT-parameter: Buffer, ReturnLength.
// INOUT-parameter: Context.
// *OPT-parameter: Buffer, ReturnLength.
func NtQueryDirectoryObject(
	DirectoryHandle Handle,
	Buffer *byte,
	Length uint32,
	ReturnSingleEntry bool,
	RestartScan bool,
	Context *uint32,
	ReturnLength *uint32,
) NtStatus {
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
func NtOpenSymbolicLinkObject(
	LinkHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
) NtStatus {
	r0, _, _ := procNtOpenSymbolicLinkObject.Call(uintptr(unsafe.Pointer(LinkHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

// OUT-parameter: ReturnedLength.
// INOUT-parameter: LinkTarget.
// *OPT-parameter: ReturnedLength.
func NtQuerySymbolicLinkObject(
	LinkHandle Handle,
	LinkTarget *UnicodeString,
	ReturnedLength *uint32,
) NtStatus {
	r0, _, _ := procNtQuerySymbolicLinkObject.Call(uintptr(LinkHandle),
		uintptr(unsafe.Pointer(LinkTarget)),
		uintptr(unsafe.Pointer(ReturnedLength)))
	return NtStatus(r0)
}

// OUT-parameter: SymbolicLinkHandle.
func NtCreateSymbolicLinkObject(
	SymbolicLinkHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
	TargetName *UnicodeString,
) NtStatus {
	r0, _, _ := procNtCreateSymbolicLinkObject.Call(uintptr(unsafe.Pointer(SymbolicLinkHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)),
		uintptr(unsafe.Pointer(TargetName)))
	return NtStatus(r0)
}

// OUT-parameter: DirectoryHandle.
func NtCreateDirectoryObject(
	DirectoryHandle *Handle,
	DesiredAccess AccessMask,
	ObjectAttributes *ObjectAttributes,
) NtStatus {
	r0, _, _ := procNtCreateDirectoryObject.Call(uintptr(unsafe.Pointer(DirectoryHandle)),
		uintptr(DesiredAccess),
		uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

// OUT-parameter: ObjectInformation, ReturnLength.
// *OPT-parameter: Handle, ReturnLength.
func NtQueryObject(
	Handle Handle,
	ObjectInformationClass ObjectInformationClass,
	ObjectInformation *byte,
	ObjectInformationLength uint32,
	ReturnLength *uint32,
) NtStatus {
	r0, _, _ := procNtQueryObject.Call(uintptr(Handle),
		uintptr(ObjectInformationClass),
		uintptr(unsafe.Pointer(ObjectInformation)),
		uintptr(ObjectInformationLength),
		uintptr(unsafe.Pointer(ReturnLength)))
	return NtStatus(r0)
}

// OUT-parameter: TargetHandle.
// *OPT-parameter: TargetProcessHandle, TargetHandle.
func NtDuplicateObject(
	SourceProcessHandle Handle,
	SourceHandle Handle,
	TargetProcessHandle Handle,
	TargetHandle *Handle,
	DesiredAccess AccessMask,
	HandleAttributes uint32,
	Options uint32,
) NtStatus {
	r0, _, _ := procNtDuplicateObject.Call(uintptr(SourceProcessHandle),
		uintptr(SourceHandle),
		uintptr(TargetProcessHandle),
		uintptr(unsafe.Pointer(TargetHandle)),
		uintptr(DesiredAccess),
		uintptr(HandleAttributes),
		uintptr(Options))
	return NtStatus(r0)
}
