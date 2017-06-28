// This file was autogenerated using go run mkcode.go -- registry.go
// DO NOT EDIT.

package ntdll

import "unsafe"

// The KeyInformationClass constants have been derived from the KEY_INFORMATION_CLASS enum definition.
type KeyInformationClass uint32

const (
	KeyBasicInformation          KeyInformationClass = 0
	KeyNodeInformation                               = 1
	KeyFullInformation                               = 2
	KeyNameInformation                               = 3
	KeyCachedInformation                             = 4
	KeyFlagsInformation                              = 5
	KeyVirtualizationInformation                     = 6
	KeyHandleTagsInformation                         = 7
	MaxKeyInfoClass                                  = 8
)

// The KeyValueInformationClass constants have been derived from the KEY_VALUE_INFORMATION_CLASS enum definition.
type KeyValueInformationClass uint32

const (
	KeyValueBasicInformation          KeyValueInformationClass = 0
	KeyValueFullInformation                                    = 1
	KeyValuePartialInformation                                 = 2
	KeyValueFullInformationAlign64                             = 3
	KeyValuePartialInformationAlign64                          = 4
	MaxKeyValueInfoClass                                       = 5
)

// The KeySetInformationClass constants have been derived from the KEY_SET_INFORMATION_CLASS enum definition.
type KeySetInformationClass uint32

const (
	KeyWriteTimeInformation         KeySetInformationClass = 0
	KeyWow64FlagsInformation                               = 1
	KeyControlFlagsInformation                             = 2
	KeySetVirtualizationInformation                        = 3
	KeySetDebugInformation                                 = 4
	KeySetHandleTagsInformation                            = 5
	MaxKeySetInfoClass                                     = 6
)

var (
	procNtCreateKey                = modntdll.NewProc("NtCreateKey")
	procNtCreateKeyTransacted      = modntdll.NewProc("NtCreateKeyTransacted")
	procNtDeleteKey                = modntdll.NewProc("NtDeleteKey")
	procNtDeleteValueKey           = modntdll.NewProc("NtDeleteValueKey")
	procNtEnumerateKey             = modntdll.NewProc("NtEnumerateKey")
	procNtEnumerateValueKey        = modntdll.NewProc("NtEnumerateValueKey")
	procNtFlushKey                 = modntdll.NewProc("NtFlushKey")
	procNtNotifyChangeKey          = modntdll.NewProc("NtNotifyChangeKey")
	procNtNotifyChangeMultipleKeys = modntdll.NewProc("NtNotifyChangeMultipleKeys")
	procNtOpenKey                  = modntdll.NewProc("NtOpenKey")
	procNtOpenKeyTransacted        = modntdll.NewProc("NtOpenKeyTransacted")
	procNtOpenKeyTransactedEx      = modntdll.NewProc("NtOpenKeyTransactedEx")
	procNtQueryKey                 = modntdll.NewProc("NtQueryKey")
	procNtQueryMultipleValueKey    = modntdll.NewProc("NtQueryMultipleValueKey")
	procNtQueryValueKey            = modntdll.NewProc("NtQueryValueKey")
	procNtRenameKey                = modntdll.NewProc("NtRenameKey")
	procNtSetInformationKey        = modntdll.NewProc("NtSetInformationKey")
	procNtSetValueKey              = modntdll.NewProc("NtSetValueKey")
)

// KeyValueEntry has been derived from the KEY_VALUE_ENTRY struct definition.
type KeyValueEntry struct {
	ValueName  *UnicodeString
	DataLength uint32
	DataOffset uint32
	Type       uint32
}

// KeyBasicInformationT has been derived from the KEY_BASIC_INFORMATION struct definition.
type KeyBasicInformationT struct {
	LastWriteTime int64
	TitleIndex    uint32
	NameLength    uint32
	Name          [1]uint16
}

// KeyNodeInformationT has been derived from the KEY_NODE_INFORMATION struct definition.
type KeyNodeInformationT struct {
	LastWriteTime int64
	TitleIndex    uint32
	ClassOffset   uint32
	ClassLength   uint32
	NameLength    uint32
	Name          [1]uint16
}

// KeyFullInformationT has been derived from the KEY_FULL_INFORMATION struct definition.
type KeyFullInformationT struct {
	LastWriteTime   int64
	TitleIndex      uint32
	ClassOffset     uint32
	ClassLength     uint32
	SubKeys         uint32
	MaxNameLen      uint32
	MaxClassLen     uint32
	Values          uint32
	MaxValueNameLen uint32
	MaxValueDataLen uint32
	Class           [1]uint16
}

// KeyNameInformationT has been derived from the KEY_NAME_INFORMATION struct definition.
type KeyNameInformationT struct {
	NameLength uint32
	Name       [1]uint16
}

// KeyCachedInformationT has been derived from the KEY_CACHED_INFORMATION struct definition.
type KeyCachedInformationT struct {
	LastWriteTime   int64
	TitleIndex      uint32
	SubKeys         uint32
	MaxNameLen      uint32
	Values          uint32
	MaxValueNameLen uint32
	MaxValueDataLen uint32
	NameLength      uint32
}

// KeyValueBasicInformationT has been derived from the KEY_VALUE_BASIC_INFORMATION struct definition.
type KeyValueBasicInformationT struct {
	TitleIndex uint32
	Type       uint32
	NameLength uint32
	Name       [1]uint16
}

// KeyValueFullInformationT has been derived from the KEY_VALUE_FULL_INFORMATION struct definition.
type KeyValueFullInformationT struct {
	TitleIndex uint32
	Type       uint32
	DataOffset uint32
	DataLength uint32
	NameLength uint32
	Name       [1]uint16
}

// KeyValuePartialInformationT has been derived from the KEY_VALUE_PARTIAL_INFORMATION struct definition.
type KeyValuePartialInformationT struct {
	TitleIndex uint32
	Type       uint32
	DataLength uint32
	Data       [1]byte
}

func NtCreateKey(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes, TitleIndex uint32, Class *UnicodeString, CreateOptions uint32, Disposition *uint32) NtStatus {
	r0, _, _ := procNtCreateKey.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)), uintptr(TitleIndex), uintptr(unsafe.Pointer(Class)), uintptr(CreateOptions), uintptr(unsafe.Pointer(Disposition)))
	return NtStatus(r0)
}

func NtCreateKeyTransacted(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes, TitleIndex uint32, Class *UnicodeString, CreateOptions uint32, TransactionHandle Handle, Disposition *uint32) NtStatus {
	r0, _, _ := procNtCreateKeyTransacted.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)), uintptr(TitleIndex), uintptr(unsafe.Pointer(Class)), uintptr(CreateOptions), uintptr(TransactionHandle), uintptr(unsafe.Pointer(Disposition)))
	return NtStatus(r0)
}

func NtDeleteKey(KeyHandle Handle) NtStatus {
	r0, _, _ := procNtDeleteKey.Call(uintptr(KeyHandle))
	return NtStatus(r0)
}

func NtDeleteValueKey(KeyHandle Handle, ValueName *UnicodeString) NtStatus {
	r0, _, _ := procNtDeleteValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueName)))
	return NtStatus(r0)
}

func NtEnumerateKey(KeyHandle Handle, Index uint32, KeyInformationClass KeyInformationClass, KeyInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtEnumerateKey.Call(uintptr(KeyHandle), uintptr(Index), uintptr(KeyInformationClass), uintptr(unsafe.Pointer(KeyInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtEnumerateValueKey(KeyHandle Handle, Index uint32, KeyValueInformationClass KeyValueInformationClass, KeyValueInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtEnumerateValueKey.Call(uintptr(KeyHandle), uintptr(Index), uintptr(KeyValueInformationClass), uintptr(unsafe.Pointer(KeyValueInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtFlushKey(KeyHandle Handle) NtStatus {
	r0, _, _ := procNtFlushKey.Call(uintptr(KeyHandle))
	return NtStatus(r0)
}

func NtNotifyChangeKey(KeyHandle Handle, Event Handle, ApcRoutine *IoApcRoutine, ApcContext *byte, IoStatusBlock *IoStatusBlock, CompletionFilter uint32, WatchTree bool, Buffer *byte, BufferSize uint32, Asynchronous bool) NtStatus {
	r0, _, _ := procNtNotifyChangeKey.Call(uintptr(KeyHandle), uintptr(Event), uintptr(unsafe.Pointer(ApcRoutine)), uintptr(unsafe.Pointer(ApcContext)), uintptr(unsafe.Pointer(IoStatusBlock)), uintptr(CompletionFilter), fromBool(WatchTree), uintptr(unsafe.Pointer(Buffer)), uintptr(BufferSize), fromBool(Asynchronous))
	return NtStatus(r0)
}

func NtNotifyChangeMultipleKeys(MasterKeyHandle Handle, Count uint32, SubordinateObjects *ObjectAttributes, Event Handle, ApcRoutine *IoApcRoutine, ApcContext *byte, IoStatusBlock *IoStatusBlock, CompletionFilter uint32, WatchTree bool, Buffer *byte, BufferSize uint32, Asynchronous bool) NtStatus {
	r0, _, _ := procNtNotifyChangeMultipleKeys.Call(uintptr(MasterKeyHandle), uintptr(Count), uintptr(unsafe.Pointer(SubordinateObjects)), uintptr(Event), uintptr(unsafe.Pointer(ApcRoutine)), uintptr(unsafe.Pointer(ApcContext)), uintptr(unsafe.Pointer(IoStatusBlock)), uintptr(CompletionFilter), fromBool(WatchTree), uintptr(unsafe.Pointer(Buffer)), uintptr(BufferSize), fromBool(Asynchronous))
	return NtStatus(r0)
}

func NtOpenKey(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes) NtStatus {
	r0, _, _ := procNtOpenKey.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)))
	return NtStatus(r0)
}

func NtOpenKeyTransacted(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes, TransactionHandle Handle) NtStatus {
	r0, _, _ := procNtOpenKeyTransacted.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)), uintptr(TransactionHandle))
	return NtStatus(r0)
}

func NtOpenKeyTransactedEx(KeyHandle *Handle, DesiredAccess AccessMask, ObjectAttributes *ObjectAttributes, OpenOptions uint32, TransactionHandle Handle) NtStatus {
	r0, _, _ := procNtOpenKeyTransactedEx.Call(uintptr(unsafe.Pointer(KeyHandle)), uintptr(DesiredAccess), uintptr(unsafe.Pointer(ObjectAttributes)), uintptr(OpenOptions), uintptr(TransactionHandle))
	return NtStatus(r0)
}

func NtQueryKey(KeyHandle Handle, KeyInformationClass KeyInformationClass, KeyInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtQueryKey.Call(uintptr(KeyHandle), uintptr(KeyInformationClass), uintptr(unsafe.Pointer(KeyInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtQueryMultipleValueKey(KeyHandle Handle, ValueEntries *KeyValueEntry, EntryCount uint32, ValueBuffer *byte, BufferLength *uint32, RequiredBufferLength *uint32) NtStatus {
	r0, _, _ := procNtQueryMultipleValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueEntries)), uintptr(EntryCount), uintptr(unsafe.Pointer(ValueBuffer)), uintptr(unsafe.Pointer(BufferLength)), uintptr(unsafe.Pointer(RequiredBufferLength)))
	return NtStatus(r0)
}

func NtQueryValueKey(KeyHandle Handle, ValueName *UnicodeString, KeyValueInformationClass KeyValueInformationClass, KeyValueInformation *byte, Length uint32, ResultLength *uint32) NtStatus {
	r0, _, _ := procNtQueryValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueName)), uintptr(KeyValueInformationClass), uintptr(unsafe.Pointer(KeyValueInformation)), uintptr(Length), uintptr(unsafe.Pointer(ResultLength)))
	return NtStatus(r0)
}

func NtRenameKey(KeyHandle Handle, NewName *UnicodeString) NtStatus {
	r0, _, _ := procNtRenameKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(NewName)))
	return NtStatus(r0)
}

func NtSetInformationKey(KeyHandle Handle, KeySetInformationClass KeySetInformationClass, KeySetInformation *byte, KeySetInformationLength uint32) NtStatus {
	r0, _, _ := procNtSetInformationKey.Call(uintptr(KeyHandle), uintptr(KeySetInformationClass), uintptr(unsafe.Pointer(KeySetInformation)), uintptr(KeySetInformationLength))
	return NtStatus(r0)
}

func NtSetValueKey(KeyHandle Handle, ValueName *UnicodeString, TitleIndex uint32, Type uint32, Data *byte, DataSize uint32) NtStatus {
	r0, _, _ := procNtSetValueKey.Call(uintptr(KeyHandle), uintptr(unsafe.Pointer(ValueName)), uintptr(TitleIndex), uintptr(Type), uintptr(unsafe.Pointer(Data)), uintptr(DataSize))
	return NtStatus(r0)
}
