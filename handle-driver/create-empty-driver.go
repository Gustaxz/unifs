package handleDriver

import "os"

func createEmptyDriver() *os.File {
	filePath := "mydriver"

	os.Remove(filePath)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sizeInBytes := 2 * 1024 * 1024
	if err := file.Truncate(int64(sizeInBytes)); err != nil {
		panic(err)
	}

	return file
}
