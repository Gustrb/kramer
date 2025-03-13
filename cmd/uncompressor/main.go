package main

import (
	"log"
	"os"
	"strings"

	"github.com/Gustrb/kramer/utils"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: uncompressor <filepath> <output>")
		return
	}

	filepath := os.Args[1]

	if strings.HasSuffix(filepath, ".gz") {
		uncompressor := utils.GzipUncompressor{}
		output, err := uncompressor.Uncompress(filepath)
		if err != nil {
			log.Fatalf("Failed to uncompress file: %s", err)
			return
		}

		log.Printf("Generated uncompressed version of %s here %s", filepath, output)
		return
	}
}
