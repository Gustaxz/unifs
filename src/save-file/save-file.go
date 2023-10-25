package saveFile

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"

	"github.com/gustaxz/unifs/src/blocks"
	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
	FAT "github.com/gustaxz/unifs/src/fat"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type File struct {
	Name [8]byte
	Ext  [3]byte
	Data []byte
}

func SaveFile(file File, f *os.File, bootSector *bootSector.BootSectorMainInfos) {
	sizeOfSector := bootSector.BytesPerSector

	fileSize := len(file.Data)
	sectorsAmount := math.Ceil(float64(fileSize) / float64(sizeOfSector))

	fmt.Println("fileSize", fileSize)

	emptyAdress, err := FAT.ListOfEmptyAdressesFAT(f, bootSector)
	check(err)

	if len(emptyAdress) < int(sectorsAmount) {
		panic("not enough space")
	}

	adresses := emptyAdress[0:int(sectorsAmount)]

	// Escrevendo na FAT
	for i, adress := range adresses {
		fmt.Println("adresses", adresses)
		fmt.Println("adress", adress)
		var nextAdress uint16

		if i == len(adresses)-1 {
			nextAdress = 65535
		} else {
			nextAdress = uint16(adresses[i+1])
		}

		buffer := make([]byte, 2)
		binary.LittleEndian.PutUint16(buffer, nextAdress)

		err := FAT.EntryAdressSectorAtFAT(buffer, adress, f, bootSector)
		check(err)
	}

	// Escrevendo no Root Directory
	entry := directoryEntry.DirectoryEntry{
		FileName:            file.Name,
		FileExtension:       file.Ext,
		FileAttributes:      0x00,
		Reserved:            [2]byte{0x00, 0x00},
		CreationTime:        [2]byte{0x00, 0x00},
		CreationDate:        [2]byte{0x00, 0x00},
		LastAccessDate:      [2]byte{0x00, 0x00},
		Ignored:             [2]byte{0x00, 0x00},
		LastWriteTime:       [2]byte{0x00, 0x00},
		LastWriteDate:       [2]byte{0x00, 0x00},
		FirstLogicalCluster: [2]byte{byte(uint16(adresses[0]))},
		FileSize:            [4]byte{byte(uint16(fileSize))},
	}

	err = directoryEntry.CreateDirectoryEntry(entry, f, bootSector)
	check(err)

	// Escrevendo no Data Region
	err = blocks.CreateSector(adresses, file.Data, f, bootSector)
	check(err)
}

func ReadFile(file File, f *os.File, bootSector *bootSector.BootSectorMainInfos) {
	fileFullName := make([]byte, 11)
	copy(fileFullName, file.Name[:])
	copy(fileFullName[8:], file.Ext[:])
	infos, err := directoryEntry.FindFileAtRootDirectoryEntry(fileFullName, f, bootSector)
	check(err)

	sectors, err := FAT.LinkedAdressesFAT(int(infos.FirstSector), f, bootSector)
	check(err)

	data, err := blocks.ReturnSector(sectors, f, bootSector)
	check(err)

	fmt.Println(string(data))
}
