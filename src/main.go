package main

import (
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
	"github.com/gustaxz/unifs/utils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// os.Remove("mydriver")

	// f, err := handleDriver.CreateEmptyDriver()
	// check(err)
	// defer f.Close()

	// err = bootSector.CreateBootSector(f)
	// check(err)

	f, err := os.Open("mydriver")
	check(err)

	bootSector := bootSector.ReadBootSector(f)
	fmt.Println(bootSector)
	f.Close()

	f, err = os.OpenFile("mydriver", os.O_RDWR, 0644)
	check(err)
	defer f.Close()

	// err = FAT.EntryAdressSectorAtFAT([]byte{0x10, 0xCC}, 0, f, &bootSector)
	// check(err)
	// err = FAT.EntryAdressSectorAtFAT([]byte{0xDD, 0x1A}, 1, f, &bootSector)
	// check(err)
	// err = FAT.EntryAdressSectorAtFAT([]byte{0xBB, 0x98}, 2, f, &bootSector)
	// check(err)
	// emptyAdress, err := FAT.ListOfEmptyAdressesFAT(f, &bootSector)
	// check(err)
	// fmt.Println(emptyAdress)

	data := directoryEntry.DirectoryEntry{
		FileName:            [8]byte(utils.StringToBytes("GUSTAVO", 8)),
		FileExtension:       [3]byte(utils.StringToBytes("TXT", 3)),
		FileAttributes:      0x00,
		Reserved:            [2]byte{0x00, 0x00},
		CreationTime:        [2]byte{0x00, 0x00},
		CreationDate:        [2]byte{0x00, 0x00},
		LastAccessDate:      [2]byte{0x00, 0x00},
		Ignored:             [2]byte{0x00, 0x00},
		LastWriteTime:       [2]byte{0x00, 0x00},
		LastWriteDate:       [2]byte{0x00, 0x00},
		FirstLogicalCluster: [2]byte{0x00, 0x00},
		FileSize:            [4]byte{0x00, 0x00, 0x00, 0x00},
	}

	err = directoryEntry.CreateDirectoryEntry(data, f, &bootSector)
	check(err)
	err = directoryEntry.CreateDirectoryEntry(data, f, &bootSector)
	check(err)
}
