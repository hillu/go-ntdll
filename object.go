package ntdll

import (
	"unsafe"
)

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS WINAPI NtOpenDirectoryObject(
  _Out_ PHANDLE            DirectoryHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS WINAPI NtQueryDirectoryObject(
  _In_      HANDLE  DirectoryHandle,
  _Out_opt_ PVOID   Buffer,
  _In_      ULONG   Length,
  _In_      BOOLEAN ReturnSingleEntry,
  _In_      BOOLEAN RestartScan,
  _Inout_   PULONG  Context,
  _Out_opt_ PULONG  ReturnLength
);
*/

/*
func:
NTSTATUS WINAPI NtOpenSymbolicLinkObject(
  _Out_ PHANDLE            LinkHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS WINAPI NtQuerySymbolicLinkObject(
  _In_      HANDLE          LinkHandle,
  _Inout_   PUNICODE_STRING LinkTarget,
  _Out_opt_ PULONG          ReturnedLength
);
*/

/*
func:
NTSTATUS WINAPI NtCreateSymbolicLinkObject (
    _Out_ PHANDLE            SymbolicLinkHandle,
    _In_  ACCESS_MASK        DesiredAccess,
    _In_  POBJECT_ATTRIBUTES ObjectAttributes,
    _In_  PUNICODE_STRING    TargetName
);
*/

/*
func:
NTSTATUS NtCreateDirectoryObject(
  _Out_ PHANDLE            DirectoryHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
type:
typedef struct _OBJECT_ATTRIBUTES {
  ULONG           Length;
  HANDLE          RootDirectory;
  PUNICODE_STRING ObjectName;
  ULONG           Attributes;
  PVOID           SecurityDescriptor;
  PVOID           SecurityQualityOfService;
} OBJECT_ATTRIBUTES, *POBJECT_ATTRIBUTES;
*/

// FIXME: PVOID -> *byte or PVOID -> uintptr?a

func NewObjectAttributes(objectName string, attr uint32, rootdir Handle, sd *byte) (oa *ObjectAttributes) {
	oa = &ObjectAttributes{
		Length:             uint32(unsafe.Sizeof(*oa)),
		RootDirectory:      rootdir,
		ObjectName:         NewUnicodeString(objectName),
		Attributes:         attr,
		SecurityDescriptor: sd,
	}
	return
}

/*
type:
typedef struct _OBJECT_DIRECTORY_INFORMATION {
    UNICODE_STRING Name;
    UNICODE_STRING TypeName;
} OBJECT_DIRECTORY_INFORMATION, *POBJECT_DIRECTORY_INFORMATION;
*/

// typedef DWORD ACCESS_MASK

type AccessMask uint32

// see winnt.h
const (
	DELETE       AccessMask = 0x00010000
	READ_CONTROL            = 0x00020000
	WRITE_DAC               = 0x00040000
	WRITE_OWNER             = 0x00080000
	SYNCHRONIZE             = 0x00100000

	STANDARD_RIGHTS_READ    = 0x00020000
	STANDARD_RIGHTS_WRITE   = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL

	STANDARD_RIGHTS_ALL = 0x001F0000

	SPECIFIC_RIGHTS_ALL    = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY = 0x01000000

	DIRECTORY_QUERY               = 0x00000001
	DIRECTORY_TRAVERSE            = 0x00000002
	DIRECTORY_CREATE_OBJECT       = 0x00000004
	DIRECTORY_CREATE_SUBDIRECTORY = 0x00000008
	DIRECTORY_ALL_ACCESS          = 0x0000000f

	KEY_QUERY_VALUE        = 0x00000001
	KEY_SET_VALUE          = 0x00000002
	KEY_CREATE_SUB_KEY     = 0x00000004
	KEY_ENUMERATE_SUB_KEYS = 0x00000008
	KEY_NOTIFY             = 0x00000010
	KEY_CREATE_LINK        = 0x00000020
	KEY_WOW64_64KEY        = 0x00000100
	KEY_WOW64_32KEY        = 0x00000200
	KEY_WOW64_RES          = 0x00000300

	KEY_READ       = ((STANDARD_RIGHTS_READ | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) &^ SYNCHRONIZE)
	KEY_WRITE      = ((STANDARD_RIGHTS_WRITE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) &^ SYNCHRONIZE)
	KEY_EXECUTE    = ((KEY_READ) &^ SYNCHRONIZE)
	KEY_ALL_ACCESS = ((STANDARD_RIGHTS_ALL | KEY_QUERY_VALUE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY | KEY_CREATE_LINK) &^ SYNCHRONIZE)
)
