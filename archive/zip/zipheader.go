package main

type FileHeader struct {
	Name               string
	CreatorVersion     uint16
	ReaderVersion      uint16
	Flags              uint16
	Method             uint16
	ModifiedTime       uint16
	ModifiedDate       uint16
	CRC32              uint32
	CompressedSize     uint32
	UncompressedSize   uint32
	CompressedSize64   uint64
	UncompressedSize64 uint64
	Extra              []byte
	ExternalAttrs      uint32
	Comment            string
}
