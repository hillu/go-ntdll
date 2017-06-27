package ntdll

//go:generate -command mkcode go run mkcode.go --
//go:generate mkcode $GOFILE

/*
type:
typedef struct _GUID {
	DWORD Data1;
	WORD  Data2;
	WORD  Data3;
	BYTE  Data4[8];
} GUID;
*/

// type Uuid Guid
