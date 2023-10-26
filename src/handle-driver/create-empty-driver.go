package handleDriver

import (
	"log"
	"os"
)

func CreateEmptyDriver(drivePath string, size int64) (*os.File, error) {

	file, err := os.Create(drivePath)
	if err != nil {
		return nil, err
	}

	if err := file.Truncate(size); err != nil {
		return nil, err
	}

	log.Println("Driver created successfully!")

	return file, nil
}
