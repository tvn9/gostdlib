package main

import "time"

type Header struct {
	Name       string
	Mode       int64
	Uid        int
	Gid        int
	Size       int64
	ModTime    time.Time
	TypeFlag   byte
	Linkname   string
	Uname      string
	Gname      string
	Devmajor   int64
	Devminor   int64
	AccessTime time.Time
	ChangeTime time.Time
}
