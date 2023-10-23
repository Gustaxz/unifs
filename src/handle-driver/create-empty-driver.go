package handleDriver

import (
	"log"
	"os"
)

func CreateEmptyDriver() (*os.File, error) {
	filePath := "mydriver"

	os.Remove(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	sizeInBytes := 2 * 1024 * 1024
	if err := file.Truncate(int64(sizeInBytes)); err != nil {
		return nil, err
	}

	log.Println("Driver created successfully!")

	return file, nil
}
