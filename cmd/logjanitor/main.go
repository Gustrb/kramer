package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Gustrb/kramer/utils"
)

const (
	NumberOfDaysToKeepUncompressed = 60
	NumberOfDaysToKeepCompressed   = 365
)

func getDateFromFileName(fname string) (time.Time, error) {
	// We follow this patter for log files:
	// logs-dd-mm-yyyy-hh:mm:ss.log
	// We need to extract the date from the file name
	if !strings.HasPrefix(fname, "logs-") {
		return time.Time{}, fmt.Errorf("invalid file name, no 'logs-' prefix")
	}

	// Remove the prefix
	fname = strings.TrimPrefix(fname, "logs-")

	// Split the file name by -
	parts := strings.Split(fname, "-")
	if len(parts) != 4 {
		return time.Time{}, fmt.Errorf("invalid file name, expected logs-dd-mm-yyyy-hh:mm:ss.log")
	}

	// only the first 3 parts are components of the date
	dateStr := strings.Join(parts[:3], "-")
	date, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format")
	}

	return date, nil
}

func shouldCompressFile(fname string) (bool, error) {
	t, err := getDateFromFileName(fname)
	if err != nil {
		return false, err
	}

	// Check if the file is older than NumberOfDaysToKeepUncompressed
	if time.Since(t).Hours()/24 > NumberOfDaysToKeepUncompressed {
		return true, nil
	}

	return false, nil
}

func shouldDeleteFile(fname string) (bool, error) {
	t, err := getDateFromFileName(fname)
	if err != nil {
		return false, err
	}

	// Check if the file is older than NumberOfDaysToKeepCompressed
	if time.Since(t).Hours()/24 > NumberOfDaysToKeepCompressed {
		return true, nil
	}

	return false, nil
}

func main() {
	if NumberOfDaysToKeepUncompressed > NumberOfDaysToKeepCompressed {
		fmt.Println("NumberOfDaysToKeepUncompressed should be less than NumberOfDaysToKeepCompressed")
		return
	}

	log.Println("Log janitor started")

	folderSizeBefore, err := utils.GetFolderSize("logs")
	if err != nil {
		log.Println("Error getting folder size:", err)
		return
	}

	// Get all the files in the logs directory
	entries, err := os.ReadDir("logs")
	if err != nil {
		log.Println("Error reading logs directory:", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			log.Printf("Skipping directory: %s\n", entry.Name())
			continue
		}

		if !strings.HasSuffix(entry.Name(), ".log") {
			log.Printf("Skipping non-log file: %s\n", entry.Name())
			continue
		}

		shoulDelete, err := shouldDeleteFile(entry.Name())
		if err != nil {
			log.Printf("Error checking if file should be deleted: %s\n", err)
			continue
		}

		if shoulDelete {
			log.Printf("Deleting old log file: %s\n", entry.Name())
			if err := os.Remove("logs/" + entry.Name()); err != nil {
				log.Printf("Error deleting file: %s\n", err)
				continue
			}

			log.Printf("Deleted file: %s\n", entry.Name())
			continue
		}

		shouldCompress, err := shouldCompressFile(entry.Name())
		if err != nil {
			log.Printf("Error checking if file should be compressed: %s\n", err)
			continue
		}

		if shouldCompress {
			log.Printf("Compressing old log file: %s\n", entry.Name())
			fname := "logs/" + entry.Name() + ".gz"

			if err := utils.CompressFile("logs/"+entry.Name(), fname); err != nil {
				log.Printf("Error compressing file: %s\n", err)
				continue
			}

			if err := os.Remove("logs/" + entry.Name()); err != nil {
				log.Printf("Error deleting file: %s\n", err)
				continue
			}

			log.Printf("Compressed file: %s\n", fname)
			continue
		}
	}

	log.Println("Log janitor finished")

	folderSizeAfter, err := utils.GetFolderSize("logs")
	if err != nil {
		log.Println("Error getting folder size:", err)
		return
	}

	log.Printf("Folder size before: %d bytes\n", folderSizeBefore)
	log.Printf("Folder size after: %d bytes\n", folderSizeAfter)
}
