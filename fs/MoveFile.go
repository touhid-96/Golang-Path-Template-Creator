package fs

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
)

// MoveFile move file path from source to destination
func MoveFile(srcPath, dstPath string) *errorwrapper.Wrapper {
	err := os.Rename(srcPath, dstPath)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(
			errtype.PathMove,
			err,
			srcPath)
}
