package main

import (
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	FAT "github.com/gustaxz/unifs/src/fat"
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

	// err = handleDriver.CreateBootSector(f)
	// check(err)

	f, err := os.Open("mydriver")
	check(err)

	bootSector := bootSector.ReadBootSector(f)
	fmt.Println(bootSector)
	f.Close()

	f, err = os.OpenFile("mydriver", os.O_RDWR, 0644)
	check(err)
	defer f.Close()

	err = FAT.EntryAdressSectorAtFAT([]byte{0x10, 0xCC}, 0, f, &bootSector)
	check(err)
	err = FAT.EntryAdressSectorAtFAT([]byte{0xDD, 0x1A}, 1, f, &bootSector)
	check(err)
	err = FAT.EntryAdressSectorAtFAT([]byte{0xBB, 0x98}, 2, f, &bootSector)
	check(err)
	emptyAdress, err := FAT.ListOfEmptyAdressesFAT(f, &bootSector)
	check(err)
	fmt.Println(emptyAdress)
}
