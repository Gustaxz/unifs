package unifs

import "os"

func DeleteDriver(driverPath string) error {
	return os.Remove(driverPath)

}
