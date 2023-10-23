package main

import (
	"fmt"
	"os"

	handleDriver "github.com/gustaxz/unifs/handle-driver"
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

	bootSector := handleDriver.ReadBootSector(f)
	fmt.Println(bootSector)
	f.Close()

	f, err = os.OpenFile("mydriver", os.O_RDWR, 0644)
	check(err)
	defer f.Close()

	err = handleDriver.EntryAdressSectorAtFAT([]byte{0x10, 0xCC}, 0, f, &bootSector)
	check(err)
	err = handleDriver.EntryAdressSectorAtFAT([]byte{0xDD, 0x1A}, 1, f, &bootSector)
	check(err)
	err = handleDriver.EntryAdressSectorAtFAT([]byte{0xBB, 0x98}, 2, f, &bootSector)
	check(err)
	v, err := handleDriver.ReadSectorFromFAT(f, 1, &bootSector)
	check(err)
	fmt.Printf("%x", v)
}
