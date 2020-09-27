package misc

import (
	"godoom/errors"
	"io"
	"os"
	"syscall"
)

func MakeDirectory(dir string) {
	err := os.Mkdir(dir, os.FileMode(0755))
	if err != nil {
		if err.(*os.PathError).Err != syscall.EEXIST {
			errors.Error("Could not make directory ", dir, ": ", err)
		}
	}
}

func FileLength(handle *os.File) (int64, error) {
	// save current pos
	savedPos, err := handle.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	// jump to end
	_, err = handle.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}
	// figure out length
	length, err := handle.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	// go back to original pos
	_, err = handle.Seek(savedPos, io.SeekStart)
	if err != nil {
		return 0, err
	}
	return length, nil
}
