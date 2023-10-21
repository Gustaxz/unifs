package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//	type BootSectorMainInfos struct {
//		OemName        [8]byte
//		BytesPerSector [2]byte
//		RootEntries    [2]byte
//		TotalSectors   [2]byte
//		SectorsPerFat  [2]byte
//		VolumeLabel    [11]byte
//		FileSystemType [8]byte
//	}
type BootSectorMainInfos struct {
	OemName        string
	BytesPerSector uint16
	RootEntries    uint16
	TotalSectors   uint16
	SectorsPerFat  uint16
	VolumeLabel    string
	FileSystemType string
}

func ReadDriver() {
	//sizeOfEachEntry := []int{3, 8, 2, 1, 2, 1, 2, 2, 1, 2, 2, 2, 2, 4, 4, 2, 1, 4, 11, 8}
	f, err := os.Open("mydriver")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 512)

	for i := 0; i < 1; i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
	}

	// infos := read.BootSector{}
	// values := reflect.TypeOf(infos)
	// for i := 0; i < values.NumField(); i++ {
	// 	for _, value := range sizeOfEachEntry {
	// 		switch i {
	// 		case 0:
	// 			copy(infos.JumpBoot[:], buf[:value])
	// 		}
	// 		buf = buf[value:]
	// 	}
	// }

	// fmt.Println(infos)

	// for _, value := range sizeOfEachEntry {
	// 	//fmt.Printf("%x\n", buf[:value])
	// 	fmt.Printf("%d", value)
	// 	fmt.Print(buf[:value])
	// 	// for i, v := range buf[:value] {
	// 	// 	fmt.Printf("%d: ", i)
	// 	// 	fmt.Printf("%c", v)
	// 	// }

	// 	buf = buf[value:]
	// }

	oemName := string(buf[0x03:0x0B])
	bytesPerSector := binary.BigEndian.Uint16(buf[0x0B:0x0D])
	rootEntries := binary.BigEndian.Uint16(buf[0x11:0x13])
	totalSectors := binary.BigEndian.Uint16(buf[0x13:0x15])
	sectorsPerFat := binary.BigEndian.Uint16(buf[0x16:0x18])
	volumeLabel := string(buf[0x2B:0x36])
	fileSystemType := string(buf[0x36:0x3E])

	bootSectorMainInfos := BootSectorMainInfos{
		OemName:        oemName,
		BytesPerSector: bytesPerSector,
		RootEntries:    rootEntries,
		TotalSectors:   totalSectors,
		SectorsPerFat:  sectorsPerFat,
		VolumeLabel:    volumeLabel,
		FileSystemType: fileSystemType,
	}

	fmt.Println(bootSectorMainInfos)

	// for i, v := range allValues {
	// 	fmt.Printf("%d: ", i)
	// 	fmt.Printf("[%d];", v)
	// }
}

func main() {
	ReadDriver()
	//read.CreateDriver()
}
