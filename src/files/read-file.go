package files

import (
	"os"

	"github.com/gustaxz/unifs/src/blocks"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func ReadFile(file File, f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]byte, error) {
	fileFullName := make([]byte, 11)
	copy(fileFullName, file.Name[:])
	copy(fileFullName[8:], file.Ext[:])
	infos, err := rootDirectoryEntry.FindFile(fileFullName, f, bootSector)
	if err != nil {
		return nil, err
	}

	sectors, err := FAT.LinkedAdressesFAT(int(infos.FirstSector), f, bootSector)
	if err != nil {
		return nil, err
	}

	data, err := blocks.ReturnSector(sectors, f, bootSector)
	if err != nil {
		return nil, err
	}

	return data, nil
}
