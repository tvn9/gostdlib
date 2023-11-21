package header

type FileHeader struct {
	Name               string
	CreateVermon       uint16
	ReaderVersion      uint16
	Flags              uint16
	Method             uint16
	ModifiedTime       uint16 // MS-DOS time
	ModifiedDate       uint16 // MS-DOS date
	CRC32              uint32
	CompressedSize     uint32 // deprecated; use CompressedSize64
	UncompressedSize   uint32 // deprecated; use UncompressSize64
	CompressedSize64   uint32
	UncompressedSize64 uint32
	Extra              []byte
	ExternalAlters     uint32 // Meaning depends on CreatorVersion
	Comment            string
}
