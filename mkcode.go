// +build ignore

// Simple tool that tries to understand simple function definitions as
// used in describing the Windwos Native API.
//
// Comments in Go files that begin with "func:" are parsed as a C
// function prototype.
//
// Comments in Go files that begin with "type:" are parsed as a C
// type definitions.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/scanner"
)

type State int

const (
	StateInit State = iota
	StateParam
	StateMember
	StateExit
)

type Direction int

const (
	DirectionUnspecified Direction = iota
	DirectionIn
	DirectionOut
	DirectionInOut
)

type FunctionParameterDefinition struct {
	Direction
	Type string
	Name string
}

type FunctionDefinition struct {
	Type   string
	Name   string
	Params []FunctionParameterDefinition
}

type StructMemberDefinition struct {
	Name string
	Type string
}

type StructDefinition struct {
	Name    string
	Members []StructMemberDefinition
}

var translation = map[string]string{
	"NTSTATUS":       "NtStatus",
	"HANDLE":         "Handle",
	"VOID":           "byte",
	"ULONG":          "uint32",
	"LONG":           "int32",
	"USHORT":         "uint16",
	"SHORT":          "int16",
	"WSTR":           "uint16",
	"BOOLEAN":        "bool",
	"LARGE_INTEGER":  "int64",
	"ULARGE_INTEGER": "uint64",
}

func translate(from string) (to string) {
	var ok bool
	if to, ok = translation[from]; ok {
		return
	}
	if from[0] == 'P' {
		return "*" + translate(from[1:])
	}
	if strings.Contains(from, "_") {
		words := strings.Split(from, "_")
		for _, w := range words {
			to = to + strings.Title(strings.ToLower(w))
		}
	}
	return
}

// ParseFunctionDefinition is a half-arsed parser for C function
// definitions as found in MSDN documentation for NTDLL member
// function.
//
// Example (adapted from
// https://msdn.microsoft.com/en-us/library/windows/hardware/ff567029(v=vs.85).aspx):
//
// NTSTATUS NtOpenSection(
//   _Out_ PHANDLE            SectionHandle,
//   _In_  ACCESS_MASK        DesiredAccess,
//   _In_  POBJECT_ATTRIBUTES ObjectAttributes
// );
func ParseFunctionDefinition(rd io.Reader) (*FunctionDefinition, error) {
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Init(rd)
	state := StateInit
	var f FunctionDefinition
	tokens := make([]string, 0)
	for {
		r := s.Scan()
		if r == scanner.EOF {
			break
		}
		switch state {
		case StateInit:
			switch r {
			case '(':
				if len(tokens) < 2 {
					return nil, fmt.Errorf("function definition needs at least a name and a type")
				}
				f.Name = tokens[0]
				for _, t := range tokens[1:] {
					switch t {
					// ignore WINAPI calling convention
					case "WINAPI":
					default:
						f.Type = t
					}
				}
				f.Type = translate(f.Type)
				tokens = tokens[0:0]
				state = StateParam
			case scanner.Ident:
				tokens = append([]string{s.TokenText()}, tokens...)
			default:
				return nil, fmt.Errorf("parse error: got %s", scanner.TokenString(r))
			}
		case StateParam:
			switch r {
			case scanner.Ident:
				tokens = append([]string{s.TokenText()}, tokens...)
			case ',', ')':
				if len(tokens) < 2 {
					return nil, fmt.Errorf("function parameter needs at least a name and a type")
				}
				p := FunctionParameterDefinition{Name: tokens[0]}
				for _, t := range tokens[1:] {
					switch t {
					case "_In_":
						p.Direction = DirectionIn
					case "_Out_", "_Out_opt_":
						p.Direction = DirectionOut
					case "_Inout_":
						p.Direction = DirectionInOut
					default:
						p.Type = t
					}
				}
				p.Type = translate(p.Type)
				f.Params = append(f.Params, p)
				tokens = tokens[0:0]
				if r == ')' {
					state = StateExit
				}
			}
		case StateExit:
			switch r {
			case ';':
				state = StateInit
				break
			}
		}
	}
	if state != StateInit {
		return nil, fmt.Errorf("parse error: wrong state %d", state)
	}
	return &f, nil
}

// ParseStructDefinition is a half-arsed parser for C struct
// definitions as found in MSDN documentation for NTDLL / Windows
// Driver types.
//
// Example:
//
// typedef struct _OBJECT_ATTRIBUTES {
//   ULONG           Length;
//   HANDLE          RootDirectory;
//   PUNICODE_STRING ObjectName;
//   ULONG           Attributes;
//   PVOID           SecurityDescriptor;
//   PVOID           SecurityQualityOfService;
// } OBJECT_ATTRIBUTES, *POBJECT_ATTRIBUTES;
func ParseStructDefinition(rd io.Reader) (*StructDefinition, error) {
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Init(rd)
	state := StateInit
	var sd StructDefinition
	var name string
	for {
		r := s.Scan()
		if r == scanner.EOF {
			break
		}
		switch state {
		case StateInit:
			e := fmt.Errorf("expecting typedef struct _X {")
			if r != scanner.Ident || s.TokenText() != "typedef" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || s.TokenText() != "struct" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || !strings.HasPrefix(s.TokenText(), "_") {
				return nil, e
			}
			name = s.TokenText()[1:]
			sd.Name = translate(name)
			if r = s.Scan(); r != '{' {
				return nil, e
			}
			state = StateMember
		case StateMember:
			if r == '}' {
				state = StateExit
				continue
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("%s: expecting type, got '%s'", name, s.TokenText())
			}
			m := StructMemberDefinition{Type: translate(s.TokenText())}
			if r = s.Scan(); r != scanner.Ident {
				return nil, fmt.Errorf("%s: expecting type / name pair, got '%s'", name, s.TokenText())
			}
			m.Name = s.TokenText()
			if r = s.Scan(); r != ';' {
				return nil, fmt.Errorf("%s: expecting semicolon after type / name pair, got '%s'", name, s.TokenText())
			}
			sd.Members = append(sd.Members, m)
		case StateExit:
			if r == ';' {
				state = StateInit
				break
			}
			var typ string
			if r == '*' {
				typ = "*"
				r = s.Scan()
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("expecting type def, got '%s'", r)
			}
			typ += s.TokenText()
			if strings.HasPrefix(typ, "*P") {
				typ = typ[2:]
			}
			if typ != name {
				return nil, fmt.Errorf("expecting type %s, got %s", name, typ)
			}
			if s.Peek() == ',' {
				s.Scan()
			}
		}
	}
	if state != StateInit {
		return nil, fmt.Errorf("parse error: wrong state %d", state)
	}
	return &sd, nil
}

func main() {
	var outfile string
	var functions []FunctionDefinition
	var structs []StructDefinition
	flag.StringVar(&outfile, "output", "", "output file")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("pass exactly one file")
	}
	infile := filepath.Clean(flag.Arg(0))
	fset := token.NewFileSet()
	if outfile == "" {
		ext := filepath.Ext(infile)
		outfile = infile[:len(infile)-len(ext)] + "_generated" + ext
	}
	f, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	file, err := parser.ParseFile(fset, "", f, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	for _, cgroup := range file.Comments {
		comment := strings.TrimSpace(cgroup.Text())
		if strings.HasPrefix(comment, "func:") {
			comment = comment[5:]
			f, err := ParseFunctionDefinition(strings.NewReader(comment))
			if err != nil {
				log.Fatal(err)
			}
			functions = append(functions, *f)
		} else if strings.HasPrefix(comment, "type:") {
			comment = comment[5:]
			s, err := ParseStructDefinition(strings.NewReader(comment))
			if err != nil {
				log.Fatal(err)
			}
			structs = append(structs, *s)
		}
	}
	f.Close()
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, `// This file was autogenerated using %s
// DO NOT EDIT.
package %s

`, strings.Join(append([]string{filepath.Base(os.Args[0])}, os.Args[1:]...), " "), file.Name)

funcs:
	for _, function := range functions {
		for _, param := range function.Params {
			if param.Type == "bool" {
				fmt.Fprintln(buf, `import "unsafe"`)
				break funcs
			}
		}
	}

	fmt.Fprintln(buf, "var (")
	for _, function := range functions {
		fmt.Fprintf(buf, `proc%[1]s = modntdll.NewProc("%[1]s")
`, function.Name)
	}
	fmt.Fprintln(buf, ")")

	for _, st := range structs {
		fmt.Fprintf(buf, "type %s struct {\n", st.Name)
		for _, m := range st.Members {
			fmt.Fprintf(buf, "%s %s\n", m.Name, m.Type)
		}
		fmt.Fprint(buf, "}\n\n")
	}

	for _, function := range functions {
		var plist, alist []string
		for _, param := range function.Params {
			alist = append(alist, fmt.Sprintf("%s %s", param.Name, param.Type))
			if param.Type == "bool" {
				plist = append(plist, fmt.Sprintf("fromBool(%s)", param.Name))
			} else if param.Type[0] == '*' {
				plist = append(plist, fmt.Sprintf("uintptr(unsafe.Pointer(%s))", param.Name))
			} else {
				plist = append(plist, fmt.Sprintf("uintptr(%s)", param.Name))
			}
			// alist = append(alist, fmt.Sprintf("p%d %s", n, param.Type))
			// plist = append(plist, "p"+strconv.Itoa(n))
		}
		fmt.Fprintf(buf, `func %[1]s(%[2]s) %[3]s {
	r0, _, _ := proc%[1]s.Call(%[4]s)
	return %[3]s(r0)
}

`, function.Name, strings.Join(alist, ", "), function.Type, strings.Join(plist, ", "))
	}

	f, err = os.Create(outfile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(string(buf.Bytes()), err)
	}
	f.Write(data)
	f.Close()
}
