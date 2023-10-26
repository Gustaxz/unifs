package files

import (
	"os"

	bootSector "github.com/gustaxz/unifs/src/boot-sector"
	"github.com/gustaxz/unifs/src/utils"
)

func CopyFrom(originPath string, fTarget *os.File, bootSector *bootSector.BootSectorMainInfos) error {
	fOrigin, err := os.OpenFile(originPath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer fOrigin.Close()

	fOriginInfo, err := fOrigin.Stat()
	if err != nil {
		return err
	}

	buf := make([]byte, fOriginInfo.Size())
	_, err = fOrigin.Read(buf)
	if err != nil {
		return err
	}

	fOriginMainInfos := File{
		Name: [8]byte(utils.StringToBytes(fOriginInfo.Name(), 8)),
		Ext:  [3]byte{},
		Data: buf,
	}

	return SaveFile(fOriginMainInfos, fTarget, bootSector)
}
