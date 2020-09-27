package wad

type FileClass struct {
	OpenFile  func(string) *File
	CloseFile func(*File)
	Read      func(*File, uint, []byte) uint64
}

type File struct {
	FileClass FileClass
	Mapped    []byte
	Length    uint
	Path      string
}
