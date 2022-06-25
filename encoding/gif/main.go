package main

type Version [6]byte

func (v Version) String() string {
	return string(v[:])
}

type Dimensions uint32

func (d Dimensions) Width() int {
	return int(d) & 0xffff
}

func (d Dimensions) Height() int {
	return int(d>>16) & 0xffff
}

type GifHeader struct {
	Version    Version
	Dimensions Dimensions
}
