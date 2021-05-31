package utils

import (
	"log"
	"os"
)

var LogFile *os.File

func WriteLogsToFile(filepath string) error {
	var err error
	// Create log file if not exist
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		_, err = os.Create(filepath)
		if err != nil {
			log.Fatalf("error creating file: %v", err)
			return err
		}
	}
	// Open log file
	LogFile, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
		return err
	}
	// Set log writer to log file insted of std
	log.SetOutput(LogFile)
	return nil
}
