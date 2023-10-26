package unifs

import (
	"bytes"
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	"github.com/gustaxz/unifs/src/utils"
)

func OpenDrive(drivePath string) (*os.File, *bootSector.BootSectorMainInfos, error) {
	if _, err := os.Stat(drivePath); err != nil {
		_, err := CreateEmptyDriver(drivePath, 2*1024*1024)
		if err != nil {
			return nil, nil, err
		}

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
	fmt.Println(bootSector)
	if !bytes.Equal(utils.StringToBytes(bootSector.OemName, 6), []byte("UNIFS")) {
		return nil, nil, err
	}

	return f, bootSector, nil
}
