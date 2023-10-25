package FAT

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
)

const fatEntrySize = 2

func ReadSectorFromFAT(f *os.File, sector int, bootSector *bootSector.BootSectorMainInfos) ([]byte, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	fmt.Println(sizeOfFat)

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
		if i == sector {
			return buf, nil
		}
	}

	return []byte{}, fmt.Errorf("sector not found")

}

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
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
		if i == sector {
			f.Seek(int64(sizeOfSector)+int64(i)*2, 0)
			err = binary.Write(f, binary.LittleEndian, adress)
			return err
		}

	}

	return fmt.Errorf("sector not found")
}

func ListOfEmptyAdressesFAT(f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]int, error) {
	var emptyAdresses []int
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	//sectorsAmount := bootSector.TotalSectors

	f.Seek(int64(sizeOfSector), 0)
	defer f.Seek(0, 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i++ {
		_, err := reader.Read(buf)

		if err != nil {
			return []int{}, err
		}

		//fmt.Printf("%s", hex.Dump(buf))
		if buf[0] == 0x00 && buf[1] == 0x00 {
			emptyAdresses = append(emptyAdresses, i)
		}
	}

	return emptyAdresses, nil
}

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
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
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
						fmt.Println(err)
					}
					break
				}
			}

			return sectors, nil
		}

	}

	return []int{}, fmt.Errorf("sector not found")
}
