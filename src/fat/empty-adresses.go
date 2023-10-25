package FAT

import (
	"bufio"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// ListOfEmptyAdressesFAT retorna uma lista de endereços vazios na FAT
func ListOfEmptyAdressesFAT(f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]int, error) {
	var emptyAdresses []int
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			return []int{}, err
		}

		if buf[0] == 0x00 && buf[1] == 0x00 {
			emptyAdresses = append(emptyAdresses, i)
		}
	}

	return emptyAdresses, nil
}
