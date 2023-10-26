package rootDirectoryEntry

import (
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

func FreePositions(f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]int, error) {
	var freePositions []int

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectorySize := int(bootSector.RootEntries) * directoryEntrySize

	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < rootDirectorySize; i++ {
		_, err := f.Read(buf)
		if err != nil {
			return nil, err
		}

		if buf[0] == 0x00 {
			freePositions = append(freePositions, i)
		}
		//fmt.Printf("%s", hex.Dump(buf))
	}

	if len(freePositions) == 0 {
		return nil, fmt.Errorf("no free positions at root directory entry")
	}

	return freePositions, nil
}
