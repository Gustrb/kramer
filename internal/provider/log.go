package provider

import (
	"fmt"
	"log"
	"os"
	"time"
)

var Logger log.Logger

func SetupLogger() error {
	// try to create a logs directory
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		if err := os.Mkdir("logs", 0755); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// create a file called logs-{timestamp}.log'
	name := fmt.Sprintf("logs-%s.log", time.Now().Format("02-01-2006 15:04:05"))
	file, err := os.Create(fmt.Sprintf("logs/%s", name))
	if err != nil {
		return err
	}

	Logger = *log.New(file, "", log.LstdFlags)

	return nil
}
