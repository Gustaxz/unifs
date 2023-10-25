package directoryEntry

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	"github.com/gustaxz/unifs/utils"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type DirectoryEntry struct {
	FileName            [8]byte
	FileExtension       [3]byte
	FileAttributes      byte
	Reserved            [2]byte
	CreationTime        [2]byte
	CreationDate        [2]byte
	LastAccessDate      [2]byte
	Ignored             [2]byte
	LastWriteTime       [2]byte
	LastWriteDate       [2]byte
	FirstLogicalCluster [2]byte
	FileSize            [4]byte
}

type DirectoryEntryMainInfos struct {
	FileName    string
	FileSize    uint32
	FirstSector uint16
}

func ListRootDirectoryEntry(f *os.File, bootSector *bootSector.BootSectorMainInfos) []DirectoryEntryMainInfos {
	var entrys []DirectoryEntryMainInfos

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectorySize := int(bootSector.RootEntries) * directoryEntrySize
	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < rootDirectorySize; i++ {
		_, err := f.Read(buf)
		check(err)

		if buf[0] != 0x00 {
			entry := DirectoryEntryMainInfos{
				FileName:    string(buf[0:11]),
				FileSize:    binary.LittleEndian.Uint32(buf[28:32]),
				FirstSector: binary.LittleEndian.Uint16(buf[26:28]),
			}
			entrys = append(entrys, entry)
		}
		//fmt.Printf("%s", hex.Dump(buf))
	}

	return entrys
}

func FreePositionsAtRootDirectoryEntry(f *os.File, bootSector *bootSector.BootSectorMainInfos) []int {
	var freePositions []int

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectorySize := int(bootSector.RootEntries) * directoryEntrySize
	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < rootDirectorySize; i++ {
		_, err := f.Read(buf)
		check(err)

		if buf[0] == 0x00 {
			freePositions = append(freePositions, i)
		}
		//fmt.Printf("%s", hex.Dump(buf))
	}

	return freePositions
}

func CreateDirectoryEntry(entry DirectoryEntry, f *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	freePositionsAtRootDirectoryEntry := FreePositionsAtRootDirectoryEntry(f, bootSector)

	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32

	if len(freePositionsAtRootDirectoryEntry) == 0 {
		panic("no free positions at root directory entry")
	}

	position := freePositionsAtRootDirectoryEntry[0]
	fmt.Println(int64(sizeOfFat+sizeOfSector) + int64(directoryEntrySize)*int64(position))
	f.Seek(int64(sizeOfFat+sizeOfSector)+int64(directoryEntrySize)*int64(position), 0)
	defer f.Seek(0, 0)

	return binary.Write(f, binary.LittleEndian, utils.EncodeToBytes(entry))
}

func FindFileAtRootDirectoryEntry(fileName []byte, f *os.File, bootSector *bootSector.BootSectorMainInfos) (*DirectoryEntryMainInfos, error) {
	sizeOfSector := bootSector.BytesPerSector
	sizeOfFat := bootSector.SectorsPerFat * sizeOfSector
	directoryEntrySize := 32
	rootDirectorySize := int(bootSector.RootEntries) * directoryEntrySize
	f.Seek(int64(sizeOfFat+sizeOfSector), 0)
	defer f.Seek(0, 0)

	buf := make([]byte, directoryEntrySize)
	for i := 0; i < rootDirectorySize; i++ {
		_, err := f.Read(buf)
		check(err)

		if bytes.Equal(buf[0:11], fileName[:]) {
			entry := DirectoryEntryMainInfos{
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
