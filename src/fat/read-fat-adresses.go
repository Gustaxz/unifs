package FAT

import (
	"bufio"
	"io"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// ReadFATAdresses lê os endereços que estão na FAT
func ReadFATAdresses(f *os.File, bootSector *bootSector.BootSectorMainInfos) ([][2]byte, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	allNonEmptyAdresses := [][2]byte{}

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return [][2]byte{}, err
			}
			break
		}

		var adress [2]byte
		copy(adress[:], buf[:2])

		if buf[0] != 0x00 {
			allNonEmptyAdresses = append(allNonEmptyAdresses, adress)
		}

	}

	return allNonEmptyAdresses, nil

}
