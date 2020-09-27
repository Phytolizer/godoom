package wad

import (
	"godoom/bytestrings"
	"godoom/diskicon"
	error2 "godoom/errors"
)

type lumpIndex int

type lumpInfo struct {
	Name     [8]byte
	WadFile  *File
	Position int
	Size     int
	Cache    []byte
	Next     lumpIndex
}

var numLumps uint = 0
var lumpHash []lumpIndex
var lumpinfo []*lumpInfo

func lumpNameHash(name string) uint {
	var result uint = 5381
	for i := 0; i < 8 && i < len(name); i += 1 {
		result = ((result << 5) ^ result) ^ uint(bytestrings.ToUpper(name[i]))
	}
	return result
}

func checkNumForName(name string) lumpIndex {
	if lumpHash != nil {
		// we have a hash table! yay
		hash := lumpNameHash(name) % numLumps
		for i := lumpHash[hash]; i != -1; i = lumpinfo[i].Next {
			if bytestrings.CompareCaseInsensitive(lumpinfo[i].Name[:], []byte(name[:8])) == 0 {
				return i
			}
		}
	} else {
		// no hash table, linear search
		for i := lumpIndex(numLumps - 1); i >= 0; i -= 1 {
			if bytestrings.CompareCaseInsensitive(lumpinfo[i].Name[:], []byte(name[:8])) == 0 {
				return i
			}
		}
	}
	return -1
}

func getNumForName(name string) lumpIndex {
	i := checkNumForName(name)
	if i < 0 {
		error2.Error("wad.getNumForName: ", name, " not found!")
	}
	return i
}

func lumpLength(lump lumpIndex) int {
	if uint(lump) >= numLumps {
		error2.Error("wad.lumpLength: ", lump, " >= numLumps")
	}
	return lumpinfo[lump].Size
}

func read(wad *File, offset uint, buffer []byte) uint64 {
	return wad.FileClass.Read(wad, offset, buffer)
}

func readLump(lump lumpIndex) []byte {
	if uint(lump) >= numLumps {
		error2.Error("wad.readLump: ", lump, " >= numLumps")
	}
	dest := make([]byte, lumpLength(lump))
	l := lumpinfo[lump]
	diskicon.BeginRead(uint64(l.Size))
	c := read(l.WadFile, uint(l.Position), dest)

	if c < uint64(l.Size) {
		error2.Error("wad.readLump: only read ", c, " of ", l.Size, " on lump ", lump)
	}

	return dest
}

func cacheLumpNum(lumpNum lumpIndex) []byte {
	if uint(lumpNum) >= numLumps {
		error2.Error("wad.cacheLumpNum: ", lumpNum, " >= numLumps")
	}

	lump := lumpinfo[lumpNum]

	var result []byte
	if lump.WadFile.Mapped != nil {
		// memory-mapped file, return from mmapped region
		result = lump.WadFile.Mapped[lump.Position:]
	} else if lump.Cache != nil {
		// already cached
		result = lump.Cache
	} else {
		// load it now
		lump.Cache = readLump(lumpNum)
		result = lump.Cache
	}
	return result
}

func CacheLumpName(name string) []byte {
	return cacheLumpNum(getNumForName(name))
}
