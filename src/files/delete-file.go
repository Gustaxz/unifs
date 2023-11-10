package files

import (
	"os"

	"github.com/gustaxz/unifs/src/blocks"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
	rootDirectoryEntry "github.com/gustaxz/unifs/src/directory-entry/root"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func DeleteFile(file File, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	fileFullName := make([]byte, 11)
	copy(fileFullName, file.Name[:])
	copy(fileFullName[8:], file.Ext[:])
	infos, pos, err := rootDirectoryEntry.FindFile(fileFullName, f, bootSector)
	if err != nil {
		return err
	}

	//Deletando entrada do Root Directory Entry
	err = rootDirectoryEntry.CreateWithPosition(directoryEntry.DirectoryEntry{}, f, bootSector, pos)
	if err != nil {
		return err
	}

	sectors, err := FAT.LinkedAdressesFAT(int(infos.FirstSector), f, bootSector)
	if err != nil {
		return err
	}

	//Deletando entradas da FAT
	for _, sector := range sectors {
		err = FAT.EntryAdressSectorAtFAT([]byte{0, 0}, sector, f, bootSector)
		if err != nil {
			return err
		}

	}

	//Deletando arquivo da Data Region
	err = blocks.CreateSector(sectors, make([]byte, bootSector.BytesPerSector), f, bootSector)
	if err != nil {
		return err
	}

	return nil
}
