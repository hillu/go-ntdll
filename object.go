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
func:
NTSTATUS NtQueryObject (
    _In_opt_ HANDLE Handle,
    _In_ OBJECT_INFORMATION_CLASS ObjectInformationClass,
    _Out_ PVOID ObjectInformation, // _Out_writes_bytes_opt_(ObjectInformationLength)
    _In_ ULONG ObjectInformationLength,
    _Out_opt_ PULONG ReturnLength
);
*/

/*
func:
NTSTATUS NtDuplicateObject (
  _In_ HANDLE SourceProcessHandle,
  _In_ HANDLE SourceHandle,
  _In_opt_ HANDLE TargetProcessHandle,
  _Out_opt_ PHANDLE TargetHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_ ULONG HandleAttributes,
  _In_ ULONG Options
);
*/

/*
func:
NTSTATUS NtCreateSection (
  _Out_ HANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
  _In_opt_ PLARGE_INTEGER MaximumSize,
  _In_ ULONG SectionPageProtection,
  _In_ ULONG AllocationAttributes,
  _In_opt_ HANDLE FileHandle
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

const (
	OBJ_INHERIT            = 0x00000002
	OBJ_PERMANENT          = 0x00000010
	OBJ_EXCLUSIVE          = 0x00000020
	OBJ_CASE_INSENSITIVE   = 0x00000040
	OBJ_OPENIF             = 0x00000080
	OBJ_OPENLINK           = 0x00000100
	OBJ_KERNEL_HANDLE      = 0x00000200
	OBJ_FORCE_ACCESS_CHECK = 0x00000400
	OBJ_VALID_ATTRIBUTES   = 0x000007F2
)

// FIXME: PVOID -> *byte or PVOID -> uintptr?a

func NewObjectAttributes(objectName string, attr uint32, rootdir Handle, sd *byte) (oa *ObjectAttributes) {
	var s *UnicodeString
	if objectName != "" {
		s = NewUnicodeString(objectName)
	}
	oa = &ObjectAttributes{
		Length:             uint32(unsafe.Sizeof(*oa)),
		RootDirectory:      rootdir,
		ObjectName:         s,
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

/*
enum:
typedef enum _OBJECT_INFORMATION_CLASS {
    ObjectBasicInformation,
    ObjectNameInformation,
    ObjectTypeInformation,
    ObjectAllInformation,
    ObjectDataInformation
} OBJECT_INFORMATION_CLASS;
*/

/*
type:
typedef struct _OBJECT_BASIC_INFORMATION {
  ULONG  Attributes;
  ACCESS_MASK  GrantedAccess;
  ULONG  HandleCount;
  ULONG  PointerCount;
  ULONG  PagedPoolUsage;
  ULONG  NonPagedPoolUsage;
  ULONG  Reserved[3];
  ULONG  NameInformationLength;
  ULONG  TypeInformationLength;
  ULONG  SecurityDescriptorLength;
  LARGE_INTEGER  CreateTime;
} OBJECT_BASIC_INFORMATION, *POBJECT_BASIC_INFORMATION;
*/

/*
type:
typedef struct _OBJECT_NAME_INFORMATION {
  UNICODE_STRING Name;
} OBJECT_NAME_INFORMATION, *POBJECT_NAME_INFORMATION;
*/

/*
type:
typedef struct _GENERIC_MAPPING {
  ACCESS_MASK GenericRead;
  ACCESS_MASK GenericWrite;
  ACCESS_MASK GenericExecute;
  ACCESS_MASK GenericAll;
} GENERIC_MAPPING;
*/

/*
type:
typedef struct _OBJECT_TYPE_INFORMATION {
  UNICODE_STRING TypeName;
  ULONG TotalNumberOfObjects;
  ULONG TotalNumberOfHandles;
  ULONG TotalPagedPoolUsage;
  ULONG TotalNonPagedPoolUsage;
  ULONG TotalNamePoolUsage;
  ULONG TotalHandleTableUsage;
  ULONG HighWaterNumberOfObjects;
  ULONG HighWaterNumberOfHandles;
  ULONG HighWaterPagedPoolUsage;
  ULONG HighWaterNonPagedPoolUsage;
  ULONG HighWaterNamePoolUsage;
  ULONG HighWaterHandleTableUsage;
  ULONG InvalidAttributes;
  GENERIC_MAPPING GenericMapping;
  ULONG ValidAccessMask;
  BOOLEAN SecurityRequired;
  BOOLEAN MaintainHandleCount;
  UCHAR TypeIndex;
  CHAR ReservedByte;
  ULONG PoolType;
  ULONG DefaultPagedPoolCharge;
  ULONG DefaultNonPagedPoolCharge;
} OBJECT_TYPE_INFORMATION, *POBJECT_TYPE_INFORMATION;
*/

/*
type:
typedef struct _OBJECT_ALL_INFORMATION {
  ULONG NumberOfObjects;
  OBJECT_TYPE_INFORMATION ObjectTypeInformation[1];
} OBJECT_ALL_INFORMATION, *POBJECT_ALL_INFORMATION;
*/

func (oi *ObjectAllInformationT) GetObjectTypeInformation() []*ObjectTypeInformationT {
	types := make([]*ObjectTypeInformationT, 0, int(oi.NumberOfObjects))
	offset := uintptr(unsafe.Pointer(&oi.ObjectTypeInformation[0]))
	for i := uintptr(0); i < uintptr(oi.NumberOfObjects); i++ {
		current := (*ObjectTypeInformationT)(unsafe.Pointer(offset))
		types = append(types, current)
		offset += unsafe.Sizeof(ObjectTypeInformationT{}) + uintptr(current.TypeName.MaximumLength)
		// padding
		if rest := offset % unsafe.Sizeof(uintptr(0)); rest != 0 {
			offset += unsafe.Sizeof(uintptr(0)) - rest
		}
	}
	return types
}

/*
type:
typedef struct _OBJECT_DATA_INFORMATION {
  BOOLEAN InheritHandle;
  BOOLEAN ProtectFromClose;
} OBJECT_DATA_INFORMATION, *POBJECT_DATA_INFORMATION;
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

    SECTION_QUERY       = 0x000001
    SECTION_MAP_WRITE   = 0x000002
    SECTION_MAP_READ    = 0x000004
    SECTION_MAP_EXECUTE = 0x000008
    SECTION_EXTEND_SIZE = 0x000010
    SECTION_ALL_ACCESS  = 0x0F001F
)
