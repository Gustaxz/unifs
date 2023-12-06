package unifs

import (
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	handleErrors "github.com/gustaxz/unifs/src/errors"
)

func OpenDrive(drivePath string) (*os.File, *bootSector.BootSectorMainInfos, error) {
	if _, err := os.Stat(drivePath); err != nil {

		err = FormatDrive(drivePath, 2*1024*1024)
		if err != nil {
			return nil, nil, err
		}
	}

	f, err := os.OpenFile(drivePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, nil, err
	}

	bootSector, err := bootSector.ReadBootSector(f)
	if err != nil {
		return nil, nil, err
	}
	if bootSector.BytesPerSector != 512 {
		return nil, nil, handleErrors.ErrFileNotFormatted
	}

	return f, bootSector, nil
}
