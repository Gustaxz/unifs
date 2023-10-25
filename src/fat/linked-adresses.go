package FAT

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

/*
LinkedAdressesFAT recebe um setor inicial, caso o valor desse setor na FAT seja 0x0000, retorna um slice vazio.
Caso o valor seja 65535 ou 0xFFFF, significa que tal arquivo só tem aquele setor.
Caso contrário, a função vai buscando os outros setores, como uma lista ligada. Assim, ela retorna um slice com os setores.
*/
func LinkedAdressesFAT(firstSector int, f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]int, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	sectorsAmount := bootSector.TotalSectors

	if firstSector < 0 || firstSector > int(sectorsAmount) {
		return []int{}, fmt.Errorf("sector out of range")
	}

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return []int{}, err
			}
			break
		}

		if i == firstSector {
			if buf[0] == 0x00 && buf[1] == 0x00 {
				return []int{}, nil
			}

			var sectors []int
			sectors = append(sectors, i)

			for {
				nextSector := binary.LittleEndian.Uint16(buf)
				if nextSector == 65535 {
					break
				}
				sectors = append(sectors, int(nextSector))
				_, err := reader.Read(buf)

				if err != nil {
					if err != io.EOF {
						return []int{}, err
					}
					break
				}
			}

			return sectors, nil
		}

	}

	return []int{}, fmt.Errorf("sector not found")
}
