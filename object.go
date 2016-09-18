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

// ACCESS_MASK
type AccessMask uint32

const (
	DIRECTORY_QUERY               AccessMask = 0x00000001 // see ddk/wdm.h
	DIRECTORY_TRAVERSE            AccessMask = 0x00000002
	DIRECTORY_CREATE_OBJECT       AccessMask = 0x00000004
	DIRECTORY_CREATE_SUBDIRECTORY AccessMask = 0x00000008
	DIRECTORY_ALL_ACCESS          AccessMask = 0x0000000f
	STANDARD_RIGHTS_READ          AccessMask = 0x00020000 // see winnt.h
)
