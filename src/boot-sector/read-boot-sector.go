package bootSector

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

// Essas são as informações mais relevantes para o funcionamento deste programa em específico
type BootSectorMainInfos struct {
	OemName        string
	BytesPerSector uint16
	RootEntries    uint16
	TotalSectors   uint16
	SectorsPerFat  uint16
	VolumeLabel    string
	FileSystemType string
}

func ReadBootSector(f *os.File) (*BootSectorMainInfos, error) {

	reader := bufio.NewReader(f)
	buf := make([]byte, 512)

	for i := 0; i < 1; i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
	}

	oemName := string(buf[0x03:0x0B])
	bytesPerSector := binary.BigEndian.Uint16(buf[0x0B:0x0D])
	rootEntries := binary.BigEndian.Uint16(buf[0x11:0x13])
	totalSectors := binary.BigEndian.Uint16(buf[0x13:0x15])
	sectorsPerFat := binary.BigEndian.Uint16(buf[0x18:0x1A])
	volumeLabel := string(buf[0x2B:0x36])
	fileSystemType := string(buf[0x36:0x3E])

	return &BootSectorMainInfos{
		OemName:        oemName,
		BytesPerSector: bytesPerSector,
		RootEntries:    rootEntries,
		TotalSectors:   totalSectors,
		SectorsPerFat:  sectorsPerFat,
		VolumeLabel:    volumeLabel,
		FileSystemType: fileSystemType,
	}, nil

}
