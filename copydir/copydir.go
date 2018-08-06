package copydir

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/akiokio/pubg-archive/copyfile"
)

// CopyDir copies dir from source to destination recursively
func CopyDir(source, destination string) error {
	// Check if source exists
	folderStats, err := os.Stat(source)
	if err != nil {
		return err
	}
	// Check is source is a directory
	if !folderStats.IsDir() {
		return errors.New("Source is not a directory")
	}

	// Check if destination exists and create if needed
	if _, err = os.Stat(destination); os.IsNotExist(err) {
		log.Printf("Creating dir: %s\n", destination)
		err = os.MkdirAll(destination, folderStats.Mode())
		if err != nil {
			return err
		}
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {
		sourceFilePath := path.Join(source, entry.Name())
		destinationFilePath := path.Join(destination, entry.Name())

		if entry.IsDir() {
			// log.Printf("Processing: %s\n", entry.Name())
			// Check if new destination exists, dont copy if exists
			if _, err := os.Stat(destinationFilePath); os.IsNotExist(err) {
				err := CopyDir(sourceFilePath, destinationFilePath)
				if err != nil {
					return err
				}
			} else {
				log.Printf("Replay already saved: %s\n", destinationFilePath)
			}
			continue
		} else {
			err := copyfile.CopyFile(sourceFilePath, destinationFilePath)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
