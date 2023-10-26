package files

import (
	"os"
	"strings"

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

	fOriginName := strings.Split(fOriginInfo.Name(), ".")

	fOriginMainInfos := File{
		Name: [8]byte(utils.StringToBytes(strings.ToUpper(fOriginName[0]), 8)),
		Ext:  [3]byte(utils.StringToBytes(strings.ToUpper(fOriginName[1]), 3)),
		Data: buf,
	}

	return SaveFile(fOriginMainInfos, fTarget, bootSector)
}
