package blocks

import (
	"encoding/binary"
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReturnSectorBlock(sectors []int, f *os.File, bootSector *bootSector.BootSectorMainInfos) []byte {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectoryEntrys := int(bootSector.RootEntries)

	offset := int(sizeOfFat) + int(sizeOfSector) + int(directoryEntrySize*rootDirectoryEntrys)

	var data []byte

	for i := 0; i < len(sectors); i++ {
		seek, err := f.Seek(int64(offset)+int64(sectors[i]*int(sizeOfSector)), 0)
		check(err)
		fmt.Println(seek, sectors[i])

		buf := make([]byte, sizeOfSector)
		_, err = f.Read(buf)
		check(err)

		data = append(data, buf...)
	}

	return data
}

func CreateSectorBlock(sectors []int, data []byte, f *os.File, bootSector *bootSector.BootSectorMainInfos) {
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
				seek, err := f.Seek(int64(offset)+int64(i*int(sizeOfSector)), 0)
				check(err)
				fmt.Println(seek, i)
				err = binary.Write(f, binary.LittleEndian, dataPieces[j])
				check(err)
				break
			}
		}

	}
}
