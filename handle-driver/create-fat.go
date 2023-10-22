package handleDriver

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gustaxz/unifs/utils"
)

const fatEntrySize = 3

func ReadSectorFromFAT(f *os.File, sector int, bootSector *BootSectorMainInfos) ([]byte, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	f.Seek(int64(sizeOfSector), 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i += 2 {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
		if i == sector {
			return utils.FirstTwelveBits(buf), nil
		} else if i == sector-1 {
			return utils.LastTwelveBits(buf), nil
		}

	}

	return []byte{}, fmt.Errorf("sector not found")

}

func EntryAdressSectorAtFAT(adress []byte, sector int, f *os.File, bootSector *BootSectorMainInfos) error {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	f.Seek(int64(sizeOfSector), 0)

	reader := bufio.NewReader(f)
	buf := make([]byte, fatEntrySize)

	for i := 0; i < (int(sizeOfFat) / fatEntrySize); i += 2 {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
		if i == sector || i == sector-1 {
			res, _ := f.Seek(int64(sizeOfSector)+int64((i+1)+fatEntrySize), 0)
			fmt.Println(res)
			err := binary.Write(f, binary.LittleEndian, adress)
			f.Seek(0, 0)
			return err
		}

	}

	f.Seek(0, 0)
	return fmt.Errorf("sector not found")
}

func CreateFileAllocationTable(f *os.File, bootSector BootSectorMainInfos) {

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector

	log.Println("Size of a Sector in bytes:", sizeOfSector)
	log.Println("Size of FAT in bytes:", sizeOfFat)

	f.Seek(int64(sizeOfSector), 0)

	//data := uint16(0xFFF)
	//byteData := byte(data >> 4)
	// err := binary.Write(f, binary.LittleEndian, utils.StringToBytes("GUSTAVO", 4))
	// if err != nil {
	// 	log.Fatal(err)
	// }

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

		if i == 0 {
			fmt.Printf("%x%x\n", utils.LastTwelveBits(buf)[0], utils.LastTwelveBits(buf)[1])

		}

	}
}
