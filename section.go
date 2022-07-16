package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtCreateSection (
  _Out_ PHANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_opt_ POBJECT_ATTRIBUTES ObjectAttributes,
  _In_opt_ PLARGE_INTEGER MaximumSize,
  _In_ ULONG SectionPageProtection,
  _In_ ULONG AllocationAttributes,
  _In_opt_ HANDLE FileHandle
);
*/

/*
func:
NTSTATUS NtOpenSection(
  _Out_ PHANDLE SectionHandle,
  _In_ ACCESS_MASK DesiredAccess,
  _In_ POBJECT_ATTRIBUTES ObjectAttributes
);
*/

/*
func:
NTSTATUS NtwMapViewOfSection(
  _In_ HANDLE SectionHandle,
  _In_ HANDLE ProcessHandle,
  _Inout_ PVOID *BaseAddress,
  _In_ ULONG_PTR ZeroBits,
  _In_ SIZE_T CommitSize,
  _Inout_opt_ PLARGE_INTEGER SectionOffset,
  _Inout_ PSIZE_T ViewSize,
  _In_ SECTION_INHERIT InheritDisposition,
  _In_ ULONG AllocationType,
  _In_ ULONG Win32Protect
);
*/

/*
enum:
typedef enum _SECTION_INHERIT {
  ViewShare = 1,
  ViewUnmap = 2
} SECTION_INHERIT;
*/
