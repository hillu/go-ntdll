package ntdll

import (
	"golang.org/x/sys/windows"
)

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
func:
NTSTATUS NtQuerySecurityObject(
  HANDLE  Handle,
  SECURITY_INFORMATION  SecurityInformation,
  PSECURITY_DESCRIPTOR  SecurityDescriptor,
  ULONG  Length,
  PULONG  LengthNeeded
);
*/

/*
func:
NTSTATUS NtSetSecurityObject(
  HANDLE Handle,
  SECURITY_INFORMATION SecurityInformation,
  PSECURITY_DESCRIPTOR SecurityDescriptor
);
*/

/*
func:
NTSTATUS NtSetSecurityObject(
	HANDLE Handle,
	SECURITY_INFORMATION SecurityInformation,
	PSECURITY_DESCRIPTOR SecurityDescriptor
);
 */

type SecurityInformationT windows.SECURITY_INFORMATION

type SecurityDescriptor windows.SECURITY_DESCRIPTOR
