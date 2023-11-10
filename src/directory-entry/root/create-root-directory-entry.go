package rootDirectoryEntry

import (
	"encoding/binary"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
	"github.com/gustaxz/unifs/src/utils"
)

func CreateWithPosition(entry directoryEntry.DirectoryEntry, f *os.File, bootSector *bootSector.BootSectorMainInfos, position int) error {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32

	f.Seek(int64(sizeOfFat+sizeOfSector)+int64(directoryEntrySize)*int64(position), 0)
	defer f.Seek(0, 0)

	return binary.Write(f, binary.LittleEndian, utils.EncodeToBytes(entry))
}

func Create(entry directoryEntry.DirectoryEntry, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	freePositionsAtRootDirectoryEntry, err := FreePositions(f, bootSector)

	if err != nil {
		return err
	}

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32

	position := freePositionsAtRootDirectoryEntry[0]
	f.Seek(int64(sizeOfFat+sizeOfSector)+int64(directoryEntrySize)*int64(position), 0)
	defer f.Seek(0, 0)

	return binary.Write(f, binary.LittleEndian, utils.EncodeToBytes(entry))
}
