package rootDirectoryEntry

import (
	"encoding/binary"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	directoryEntry "github.com/gustaxz/unifs/src/directory-entry"
)

func List(f *os.File, bootSector *bootSector.BootSectorMainInfos) ([]directoryEntry.DirectoryEntryMainInfos, error) {
	var entrys []directoryEntry.DirectoryEntryMainInfos

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32

	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < int(bootSector.RootEntries); i++ {
		_, err := f.Read(buf)
		if err != nil {
			return nil, err
		}

		if buf[0] != 0x00 {
			entry := directoryEntry.DirectoryEntryMainInfos{
				FileName:    string(buf[0:11]),
				FileSize:    binary.LittleEndian.Uint32(buf[28:32]),
				FirstSector: binary.LittleEndian.Uint16(buf[26:28]),
			}
			entrys = append(entrys, entry)
		}
	}

	return entrys, nil
}
