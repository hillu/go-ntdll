// lsobj is a console-based WinObj-like tool
// See https://technet.microsoft.com/en-us/sysinternals/winobj
package main

import (
	"errors"
	"fmt"
  "os"
	"unsafe"

	"github.com/hillu/go-ntdll"
)

var SkipDir = errors.New("skip this directory")

type WalkFunc func(string, string) error

func Walk(entry string, fn WalkFunc) error {
	var h ntdll.Handle
	if st := ntdll.NtOpenDirectoryObject(&h, ntdll.STANDARD_RIGHTS_READ|ntdll.DIRECTORY_QUERY,
		ntdll.NewObjectAttributes(entry, 0, 0, nil),
	); st != 0 {
		return st.Error()
	}
	defer ntdll.NtClose(h)
	var context uint32
	for {
		var buf [32768]byte
		var length uint32
		switch st := ntdll.NtQueryDirectoryObject(h, &buf[0], uint32(len(buf)), true, context == 0, &context, &length); st {
		case 0:
		case 0x8000001a: // STATUS_NO_MORE_ENTRIES
			return nil
		default:
			return st.Error()
		}
		odi := (*ntdll.ObjectDirectoryInformationT)(unsafe.Pointer(&buf[0]))
		var path string
		if entry == `\` {
			path = `\` + odi.Name.String()
		} else {
			path = entry + `\` + odi.Name.String()
		}
		switch typ := odi.TypeName.String(); typ {
		case "Directory":
			if err := Walk(path, fn); err != nil {
				return err
			}
		default:
			switch err := fn(path, typ); err {
			case SkipDir:
				return nil
			case nil:
				continue
			default:
				return err
			}
		}
	}
	return nil
}

func main() {
	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "\\"
	}
	Walk(path, func(path, typ string) error {
		switch typ {
		case "SymbolicLink":
			var h ntdll.Handle
			if st := ntdll.NtOpenSymbolicLinkObject(&h, ntdll.STANDARD_RIGHTS_READ|ntdll.DIRECTORY_QUERY,
				ntdll.NewObjectAttributes(path, 0, 0, nil),
			); st != 0 {
				fmt.Printf("%s %s -> <NtOpenSymbolicLinkObject error %08x>\n", path, typ, st)
				return nil
			}
			defer ntdll.NtClose(h)
			target := ntdll.NewEmptyUnicodeString(1024)
			var length uint32
			if st := ntdll.NtQuerySymbolicLinkObject(h, target, &length); st != 0 {
				fmt.Printf("%s %s -> <NtQuerySymbolicLinkObject error %08x>\n", path, typ, st)
				return nil
			}
			fmt.Printf("%s %s -> %s\n", path, typ, target)
		default:
			fmt.Printf("%s %s\n", path, typ)
		}
		return nil
	})
}
