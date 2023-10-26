package rootDirectoryEntry

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
)

func FindFile(fileName []byte, f *os.File, bootSector *bootSector.BootSectorMainInfos) (*directoryEntry.DirectoryEntryMainInfos, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectorySize := int(bootSector.RootEntries) * directoryEntrySize
	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < rootDirectorySize; i++ {
		_, err := f.Read(buf)
		if err != nil {
			return nil, err
		}

		if bytes.Equal(buf[0:11], fileName[:]) {
			entry := directoryEntry.DirectoryEntryMainInfos{
				FileName:    string(buf[0:11]),
				FileSize:    binary.LittleEndian.Uint32(buf[28:32]),
				FirstSector: binary.LittleEndian.Uint16(buf[26:28]),
			}
			return &entry, nil

		}
		//fmt.Printf("%s", hex.Dump(buf))
	}

	return nil, fmt.Errorf("file not found")
}
