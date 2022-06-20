package main

import (
	"archive/tar"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

var HeaderTemplate = `tar header
Name: {{.Name}}
Mode: {{.Mode | printf "%o" }}
UID: {{.Uid}}
GID: {{.Gid}}
Size: {{.Size}}
ModTime: {{.ModTime}}
Typeflag: {{.Typeflag | printf "%q" }}
Linkname: {{.Linkname}}
Uname: {{.Uname}}
Gname: {{.Gname}}
Devmajor: {{.Devmajor}}
Devminor: {{.Devminor}}
AccessTime {{.AccessTime}}
ChangeTime: {{.ChangeTime}}
`
var HeaderTmpl *template.Template

func init() {
	t := template.New("header")
	HeaderTmpl = template.Must(t.Parse(HeaderTemplate))
}

func printHeader(hdr *tar.Header) {
	HeaderTmpl.Execute(os.Stdout, hdr)
}

func printContents(tr io.Reader, size int64) {
	contents := make([]byte, size)
	read, err := io.ReadFull(tr, contents)
	if err != nil {
		log.Fatalf("failed reading tar entry: %s", err)
	}
	if int64(read) != size {
		log.Fatalf("read %d bytes but expected to read %d", read, size)
	}
	fmt.Fprintf(os.Stdout, "Contents: \n\n%s", contents)
}
