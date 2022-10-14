package deletepaths

import (
	"os"

	"gitlab.com/evatix-go/errorwrapper"
	"gitlab.com/evatix-go/errorwrapper/errnew"
	"gitlab.com/evatix-go/errorwrapper/errtype"
	"gitlab.com/evatix-go/pathhelper/internal/fsinternal"
)

func SingleOnExistIf(
	isRemove bool,
	location string,
) *errorwrapper.Wrapper {
	if !isRemove {
		return nil
	}

	if !fsinternal.IsPathExists(location) {
		return nil
	}

	err := os.Remove(location)

	if err == nil {
		return nil
	}

	return errnew.
		Path.
		Error(errtype.DeletePathFailed, err, location)
}
