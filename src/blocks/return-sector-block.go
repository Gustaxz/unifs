package blocks

import (
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// Lê um setor no Data Region do disco
func ReturnSector(sectors []int, f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]byte, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectoryEntrys := int(bootSector.RootEntries)

	offset := int(sizeOfFat) + int(sizeOfSector) + int(directoryEntrySize*rootDirectoryEntrys)

	var data []byte

	for i := 0; i < len(sectors); i++ {
		_, err := f.Seek(int64(offset)+int64(sectors[i]*int(sizeOfSector)), 0)
		if err != nil {
			return nil, err
		}

		buf := make([]byte, sizeOfSector)
		_, err = f.Read(buf)
		if err != nil {
			return nil, err
		}
		data = append(data, buf...)
	}

	return data, nil
}
