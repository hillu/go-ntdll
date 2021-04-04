package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtQueryInformationProcess(
  HANDLE           ProcessHandle,
  PROCESSINFOCLASS ProcessInformationClass,
  PVOID            ProcessInformation,
  ULONG            ProcessInformationLength,
  PULONG           ReturnLength
);
*/

/*
enum:
typedef enum _PROCESSINFOCLASS {
  ProcessBasicInformation,
  ProcessQuotaLimits,
  ProcessIoCounters,
  ProcessVmCounters,
  ProcessTimes,
  ProcessBasePriority,
  ProcessRaisePriority,
  ProcessDebugPort,
  ProcessExceptionPort,
  ProcessAccessToken,
  ProcessLdtInformation,
  ProcessLdtSize,
  ProcessDefaultHardErrorMode,
  ProcessIoPortHandlers,
  ProcessPooledUsageAndLimits,
  ProcessWorkingSetWatch,
  ProcessUserModeIOPL,
  ProcessEnableAlignmentFaultFixup,
  ProcessPriorityClass,
  ProcessWx86Information,
  ProcessHandleCount,
  ProcessAffinityMask,
  ProcessPriorityBoost,
  ProcessDeviceMap,
  ProcessSessionInformation,
  ProcessForegroundInformation,
  ProcessWow64Information,
  ProcessImageFileName,
  ProcessLUIDDeviceMapsEnabled,
  ProcessBreakOnTermination,
  ProcessDebugObjectHandle,
  ProcessDebugFlags,
  ProcessHandleTracing,
  ProcessIoPriority,
  ProcessExecuteFlags,
  ProcessTlsInformation,
  ProcessCookie,
  ProcessImageInformation,
  ProcessCycleTime,
  ProcessPagePriority,
  ProcessInstrumentationCallback,
  ProcessThreadStackAllocation,
  ProcessWorkingSetWatchEx,
  ProcessImageFileNameWin32,
  ProcessImageFileMapping,
  ProcessAffinityUpdateMode,
  ProcessMemoryAllocationMode,
  ProcessGroupInformation,
  ProcessTokenVirtualizationEnabled,
  ProcessConsoleHostProcess,
  ProcessWindowInformation,
  MaxProcessInfoClass
} PROCESSINFOCLASS;
*/

/*
func:
NTSTATUS NtQueryInformationThread(
  HANDLE          ThreadHandle,
  THREADINFOCLASS ThreadInformationClass,
  PVOID           ThreadInformation,
  ULONG           ThreadInformationLength,
  PULONG          ReturnLength
);
*/

/*
func:
NTSTATUS NtSetInformationThread(
  HANDLE          ThreadHandle,
  THREADINFOCLASS ThreadInformationClass,
  PVOID           ThreadInformation,
  ULONG           ThreadInformationLength
);
*/

/*
func:
NTSTATUS NtSetInformationProcess (
  HANDLE           ProcessHandle,
  PROCESSINFOCLASS ProcessInformationClass,
  PVOID            ProcessInformation,
  ULONG            ProcessInformationLength
);
*/

/*
enum:
typedef enum _THREADINFOCLASS {
  ThreadBasicInformation,
  ThreadTimes,
  ThreadPriority,
  ThreadBasePriority,
  ThreadAffinityMask,
  ThreadImpersonationToken,
  ThreadDescriptorTableEntry,
  ThreadEnableAlignmentFaultFixup,
  ThreadEventPair_Reusable,
  ThreadQuerySetWin32StartAddress,
  ThreadZeroTlsCell,
  ThreadPerformanceCount,
  ThreadAmILastThread,
  ThreadIdealProcessor,
  ThreadPriorityBoost,
  ThreadSetTlsArrayAddress,
  ThreadIsIoPending,
  ThreadHideFromDebugger,
  ThreadBreakOnTermination,
  ThreadSwitchLegacyState,
  ThreadIsTerminated,
  ThreadLastSystemCall,
  ThreadIoPriority,
  ThreadCycleTime,
  ThreadPagePriority,
  ThreadActualBasePriority,
  ThreadTebInformation,
  ThreadCSwitchMon,
  ThreadCSwitchPmu,
  ThreadWow64Context,
  ThreadGroupInformation,
  ThreadUmsInformation,
  ThreadCounterProfiling,
  ThreadIdealProcessorEx,
  MaxThreadInfoClass
} THREADINFOCLASS;
*/
