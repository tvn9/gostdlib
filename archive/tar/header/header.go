// Tar header file
package header

import "time"

type Header struct {
	Name       string    // name of the header file entry
	Mode       int64     // permission and mode bits
	Uid        int       // owner user id
	Gid        int       // owner group id
	Size       int64     // length in bytes
	ModTime    time.Time // modified time
	Typeflag   byte      // type of header entry
	Linkname   string    // link target name
	Uname      string    // owner user name
	Gname      string    // owner group name
	Devmajor   int64     // block device major number of character
	Devminor   int64     // block device minor number of character
	AccessTime time.Time // access time
	ChangeTime time.Time // status change time
}
