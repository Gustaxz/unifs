package files

import (
	"os"

	"github.com/gustaxz/unifs/src/blocks"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func DeleteFile(file File, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	fileFullName := make([]byte, 11)
	copy(fileFullName, file.Name[:])
	copy(fileFullName[8:], file.Ext[:])
	infos, _, err := rootDirectoryEntry.FindFile(fileFullName, f, bootSector)
	if err != nil {
		return err
	}

	// TODO - Delete file from root directory entry

	sectors, err := FAT.LinkedAdressesFAT(int(infos.FirstSector), f, bootSector)
	if err != nil {
		return err
	}

	//Deleting file from FAT
	for _, sector := range sectors {
		err = FAT.EntryAdressSectorAtFAT([]byte{0, 0}, sector, f, bootSector)
		if err != nil {
			return err
		}

	}

	//Deleting file from Data Region
	err = blocks.CreateSector(sectors, make([]byte, bootSector.BytesPerSector), f, bootSector)
	if err != nil {
		return err
	}

	return nil
}
