package blocks

import (
	"encoding/binary"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// Cria um setor no Data Region do disco
func CreateSector(sectors []int, data []byte, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectoryEntrys := int(bootSector.RootEntries)

	offset := int(sizeOfFat) + int(sizeOfSector) + int(directoryEntrySize*rootDirectoryEntrys)

	var dataPieces [][]byte

	for i := 0; i < len(data); i += int(sizeOfSector) {
		end := i + int(sizeOfSector)
		if end > len(data) {
			end = len(data)
		}

		dataPieces = append(dataPieces, data[i:end])
	}

	availableSectors := int(bootSector.TotalSectors) - int(bootSector.SectorsPerFat*2) - int(rootDirectoryEntrys)

	for i := 0; i < availableSectors; i++ {

		for j, sector := range sectors {
			if sector == i {
				_, err := f.Seek(int64(offset)+int64(i*int(sizeOfSector)), 0)
				if err != nil {
					return err
				}
				err = binary.Write(f, binary.LittleEndian, dataPieces[j])
				if err != nil {
					return err
				}
				break
			}
		}

	}

	return nil
}
