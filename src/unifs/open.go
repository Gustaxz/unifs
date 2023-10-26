package unifs

import (
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

func OpenDrive(drivePath string) (*os.File, *bootSector.BootSectorMainInfos, error) {
	if _, err := os.Stat(drivePath); err != nil {
		_, err := CreateEmptyDriver(drivePath, 2*1024*1024)
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

	return f, bootSector, nil
}
