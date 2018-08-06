package copyfile

import (
	"io"
	"os"
)

// CopyFile copies dir from source to destination recursively
func CopyFile(source, destination string) error {
	// fmt.Printf("Copying file: %s", source)
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	sourceStats, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.Chmod(destination, sourceStats.Mode())
	if err != nil {
		return err
	}

	return nil
}
