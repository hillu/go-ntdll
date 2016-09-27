package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtCreateKey(
  _Out_      PHANDLE            KeyHandle,
  _In_       ACCESS_MASK        DesiredAccess,
  _In_       POBJECT_ATTRIBUTES ObjectAttributes,
  _Reserved_ ULONG              TitleIndex,
  _In_opt_   PUNICODE_STRING    Class,
  _In_       ULONG              CreateOptions,
  _Out_opt_  PULONG             Disposition
);
*/

/*
func:
NTSTATUS NtDeleteKey(
  _In_ HANDLE KeyHandle
);
*/

/*
func:
NTSTATUS NtEnumerateKey(
  _In_      HANDLE                KeyHandle,
  _In_      ULONG                 Index,
  _In_      KEY_INFORMATION_CLASS KeyInformationClass,
  _Out_opt_ PVOID                 KeyInformation,
  _In_      ULONG                 Length,
  _Out_     PULONG                ResultLength
);
*/

/*
func:
NTSTATUS NtEnumerateValueKey(
  _In_      HANDLE                      KeyHandle,
  _In_      ULONG                       Index,
  _In_      KEY_VALUE_INFORMATION_CLASS KeyValueInformationClass,
  _Out_opt_ PVOID                       KeyValueInformation,
  _In_      ULONG                       Length,
  _Out_     PULONG                      ResultLength
);
*/

/*
func:
NTSTATUS NtFlushKey(
  _In_ HANDLE KeyHandle
);
*/

/*
func:
NTSTATUS NtNotifyChangeKey(
  _In_      HANDLE           KeyHandle,
  _In_opt_  HANDLE           Event,
  _In_opt_  PIO_APC_ROUTINE  ApcRoutine,
  _In_opt_  PVOID            ApcContext,
  _Out_     PIO_STATUS_BLOCK IoStatusBlock,
  _In_      ULONG            CompletionFilter,
  _In_      BOOLEAN          WatchTree,
  _Out_opt_ PVOID            Buffer,
  _In_      ULONG            BufferSize,
  _In_      BOOLEAN          Asynchronous
);
*/

/*
func:
NTSTATUS NtOpenKey(
  _Out_ PHANDLE            KeyHandle,
  _In_  ACCESS_MASK        DesiredAccess,
  _In_  POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS NtQueryValueKey(
  _In_      HANDLE                      KeyHandle,
  _In_      PUNICODE_STRING             ValueName,
  _In_      KEY_VALUE_INFORMATION_CLASS KeyValueInformationClass,
  _Out_opt_ PVOID                       KeyValueInformation,
  _In_      ULONG                       Length,
  _Out_     PULONG                      ResultLength
);
*/

/*
func:
NTSTATUS NtSetValueKey(
  _In_     HANDLE          KeyHandle,
  _In_     PUNICODE_STRING ValueName,
  _In_opt_ ULONG           TitleIndex,
  _In_     ULONG           Type,
  _In_opt_ PVOID           Data,
  _In_     ULONG           DataSize
);
*/

/*
enum:
typedef enum _KEY_INFORMATION_CLASS {
  KeyBasicInformation           = 0,
  KeyNodeInformation            = 1,
  KeyFullInformation            = 2,
  KeyNameInformation            = 3,
  KeyCachedInformation          = 4,
  KeyFlagsInformation           = 5,
  KeyVirtualizationInformation  = 6,
  KeyHandleTagsInformation      = 7,
  MaxKeyInfoClass               = 8
} KEY_INFORMATION_CLASS;
*/

/*
enum:
typedef enum _KEY_VALUE_INFORMATION_CLASS {
  KeyValueBasicInformation           = 0,
  KeyValueFullInformation,
  KeyValuePartialInformation,
  KeyValueFullInformationAlign64,
  KeyValuePartialInformationAlign64,
  MaxKeyValueInfoClass
} KEY_VALUE_INFORMATION_CLASS;
*/

/*
typedef struct _IO_STATUS_BLOCK {
  union {
    NTSTATUS Status;
    PVOID    Pointer;
  };
  ULONG_PTR Information;
} IO_STATUS_BLOCK, *PIO_STATUS_BLOCK;
*/

type IoStatusBlock struct {
	StatusPointer uintptr
	Invormation   uintptr
}

/*
typedef VOID (NTAPI *PIO_APC_ROUTINE)(PVOID ApcContext,PIO_STATUS_BLOCK IoStatusBlock,ULONG Reserved);
*/

type IoApcRoutine uintptr
