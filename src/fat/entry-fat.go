package FAT

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

// O tamanho de um setor na FAT é de 2 bytes
const fatEntrySize = 2

// Adiciona endereços a tabela FAT
func EntryAdressSectorAtFAT(adress []byte, sector int, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	sectorsAmount := bootSector.TotalSectors

	if len(adress) != 2 {
		return fmt.Errorf("adress must be 2 bytes")
	}

	if sector < 0 || sector > int(sectorsAmount) {
		return fmt.Errorf("sector out of range")
	}

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		if i == sector {
			f.Seek(int64(sizeOfSector)+int64(i)*2, 0)
			err = binary.Write(f, binary.LittleEndian, adress)
			return err
		}

	}

	return fmt.Errorf("sector not found")
}
