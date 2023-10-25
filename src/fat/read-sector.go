package FAT

import (
	"bufio"
	"fmt"
	"io"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// ReadSectorFromFAT lê os setores presentes na FAT, que tem tamanho de 2 bytes
func ReadSectorFromFAT(f *os.File, sector int, bootSector *bootSector.BootSectorMainInfos) ([]byte, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return []byte{}, fmt.Errorf("sector not found")
			}
			break
		}

		if i == sector {
			return buf, nil
		}
	}

	return []byte{}, fmt.Errorf("sector not found")

}
