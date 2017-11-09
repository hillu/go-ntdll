//+build ignore

// simple crawler to extract Status Codes from MSDN.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const URL = "https://msdn.microsoft.com/en-us/library/cc704588.aspx"

type status struct {
	val  uint64
	name string
	desc string
}

type StatusList []status

func (l StatusList) Len() int           { return len(l) }
func (l StatusList) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l StatusList) Less(i, j int) bool { return l[i].val < l[j].val }

var list StatusList

func main() {
	doc, err := goquery.NewDocument(URL)
	if err != nil {
		log.Fatal(err)
	}

	/*
		<table Responsive="true" summary="table">
		 <tr class="thead" Responsive="true"><th scope="col">
		   <p>Return value/code</p>
		   </th><th scope="col">
		   <p>Description</p>
		   </th></tr>
		 <tr><td data-th="&#xA;   Return value/code&#xA;   ">
		  <p>0x00000000</p>
		  <p>STATUS_SUCCESS</p>
		  </td><td data-th="&#xA;   Description&#xA;   ">
		  <p>The operation completed successfully. </p>
		  </td></tr>
		 <tr><td data-th="&#xA;   Return value/code&#xA;   ">
		  <p>0x00000000</p>
		  <p>STATUS_WAIT_0</p>
		  </td><td data-th="&#xA;   Description&#xA;   ">
		  <p>The caller specified WaitAny for WaitType and one of
		  the dispatcher objects in the Object array has been set to the signaled
		  state.</p>
		  </td></tr>
	*/

	var n2s = map[uint64]string{}
	var spaces = regexp.MustCompile("  *")
	doc.Find("table tr").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return // skip header
		}
		st := status{}

		td1 := s.Find("td")
		if td1 == nil {
			log.Println("entry ", i, "empty td")
			return
		}

		p := td1.Find("p").First()
		v := strings.TrimSpace(p.Text())
		if _, err = fmt.Sscanf(v, "0x%X", &st.val); err != nil {
			log.Println("entry ", i, err)
			return
		}
		st.name = p.Next().Text()
		n2s[st.val] = st.name // Any duplicate gets overwritten!
		desc := strings.Replace(td1.Next().Find("p").Text(), "\n", "", -1)
		st.desc = spaces.ReplaceAllLiteralString(desc, " ")
		list = append(list, st)
	})

	// Sort the list by status code.
	sort.Sort(list)

	// Write the header
	code := &bytes.Buffer{}
	code.WriteString("package ntdll\n")
	code.WriteString("// generated by github.com/oec/go-ntdll/statuscrawler; DO NOT EDIT\n")
	fmt.Fprintf(code, "// Extracted from %s on %s\n\n", URL, time.Now().Format(time.RFC822))

	// Write the constants
	code.WriteString("const (\n")
	for i, s := range list {
		if i == 0 {
			fmt.Fprintf(code, "%s NtStatus = 0x%08X // %s\n", s.name, s.val, s.desc)
		} else {
			fmt.Fprintf(code, "%s = 0x%08X // %s\n", s.name, s.val, s.desc)
		}
	}
	code.WriteString(")\n\n")

	n2s[0x00000000] = "STATUS_SUCCESS" // Fix the correspondence

	// Write the reverse-lookup-table.
	code.WriteString("var ntStatus2str = map[NtStatus]string{\n")
	for _, status := range list {
		if short, ok := n2s[status.val]; ok {
			fmt.Fprintf(code, "0x%08X : \"%s\",\n", status.val, short)
			delete(n2s, status.val) // Avoid duplicates.
		}
	}
	code.WriteString("}\n\n")

	// Write the Stringer method for NtStatus
	code.WriteString(`
func (status NtStatus) String() string {
	return ntStatus2str[status]
}
	`)

	out, err := format.Source(code.Bytes())
	if err != nil {
		log.Println(err)
		log.Println(code.String())
	}

	if f, err := os.Create("ntstatus_generated.go"); err != nil {
		log.Fatal(err)
	} else if n, err := f.Write(out); err != nil {
		log.Fatal(err)
	} else if n != len(out) {
		log.Fatal("output size mismatch")
	} else {
		f.Close()
	}
}
