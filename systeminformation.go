//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

// Information taken from
// - http://www.ntinternals.net/: UserMode / NTDLL / System Information
// - http://www.geoffchappell.com/studies/windows/km/index.htm

package ntdll

import (
	"reflect"
	"unsafe"
)

/*
func:
NTSTATUS NtQuerySystemInformation(
  _In_  SYSTEM_INFORMATION_CLASS SystemInformationClass,
  _Out_ PVOID                    SystemInformation,
  _In_  ULONG                    SystemInformationLength,
  _Out_ PULONG                   ReturnLength
);
*/

/*
func:
NTSTATUS NtSetSystemInformation(
  _In_ SYSTEM_INFORMATION_CLASS SystemInformationClass,
  _In_ PVOID                    SystemInformation,
  _In_ ULONG                    SystemInformationLength
);
*/

/*
enum:
typedef enum _SYSTEM_INFORMATION_CLASS {
    SystemBasicInformation,
    SystemProcessorInformation,
    SystemPerformanceInformation,
    SystemTimeOfDayInformation,
    SystemPathInformation,
    SystemProcessInformation,
    SystemCallCountInformation,
    SystemDeviceInformation,
    SystemProcessorPerformanceInformation,
    SystemFlagsInformation,
    SystemCallTimeInformation,
    SystemModuleInformation,
    SystemLocksInformation,
    SystemStackTraceInformation,
    SystemPagedPoolInformation,
    SystemNonPagedPoolInformation,
    SystemHandleInformation,
    SystemObjectInformation,
    SystemPageFileInformation,
    SystemVdmInstemulInformation,
    SystemVdmBopInformation,
    SystemFileCacheInformation,
    SystemPoolTagInformation,
    SystemInterruptInformation,
    SystemDpcBehaviorInformation,
    SystemFullMemoryInformation,
    SystemLoadGdiDriverInformation,
    SystemUnloadGdiDriverInformation,
    SystemTimeAdjustmentInformation,
    SystemSummaryMemoryInformation,
    SystemNextEventIdInformation,
    SystemEventIdsInformation,
    SystemCrashDumpInformation,
    SystemExceptionInformation,
    SystemCrashDumpStateInformation,
    SystemKernelDebuggerInformation,
    SystemContextSwitchInformation,
    SystemRegistryQuotaInformation,
    SystemExtendServiceTableInformation,
    SystemPrioritySeperation,
    SystemPlugPlayBusInformation,
    SystemDockInformation,
    SystemPowerInformation,
    SystemProcessorSpeedInformation,
    SystemCurrentTimeZoneInformation,
    SystemLookasideInformation
} SYSTEM_INFORMATION_CLASS, *PSYSTEM_INFORMATION_CLASS;
*/

/*
type:
typedef struct _SYSTEM_MODULE {
  PVOID                Section;
  PVOID                MappedBase;
  PVOID                ImageBase;
  ULONG                ImageSize;
  ULONG                Flags;
  WORD                 LoadOrderIndex;
  WORD                 InitOrderIndex;
  WORD                 LoadCount;
  WORD                 NameOffset;
  BYTE                 Name[256];
} SYSTEM_MODULE, *PSYSTEM_MODULE;
*/

/*
type:
typedef struct _SYSTEM_MODULE_INFORMATION {
  ULONG                ModulesCount;
  SYSTEM_MODULE        Modules[ANYSIZE_ARRAY];
} SYSTEM_MODULE_INFORMATION, *PSYSTEM_MODULE_INFORMATION;
*/

/*
type:
typedef struct _SYSTEM_HANDLE_TABLE_ENTRY_INFO {
  ULONG                ProcessId;
  BYTE                 ObjectTypeNumber;
  BYTE                 Flags;
  USHORT               Handle;
  PVOID                Object;
  BYTE          GrantedAccess;
} SYSTEM_HANDLE_TABLE_ENTRY_INFO, *PSYSTEM_HANDLE_TABLE_ENTRY_INFO;
*/

/*
type:
typedef struct _SYSTEM_HANDLE_INFORMATION {
  ULONG                HandleCount;
  SYSTEM_HANDLE_TABLE_ENTRY_INFO Handles [ANYSIZE_ARRAY];
} SYSTEM_HANDLE_INFORMATION, *PSYSTEM_HANDLE_INFORMATION;
*/

func (mi *SystemModuleInformationT) GetModules() []SystemModule {
	modules := make([]SystemModule, int(mi.ModulesCount))
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&modules))
	hdr.Data = uintptr(unsafe.Pointer(&mi.Modules[0]))
	return modules
}

func (mi *SystemHandleInformationT) GetHandles() []SystemHandleTableEntryInfo {
	handles := make([]SystemHandleTableEntryInfo, int(mi.HandleCount))
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&handles))
	hdr.Data = uintptr(unsafe.Pointer(&mi.Handles[0]))
	return handles
}
