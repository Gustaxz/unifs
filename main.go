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
	handleDriver.CreateFileAllocationTable(f, bootSector)
}
