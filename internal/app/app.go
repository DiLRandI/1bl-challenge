package app

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Execute(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to read line: %v", err)

			break
		}

		lines := strings.Split(line, ";")
		if len(lines) < 2 {
			log.Fatalf("invalid line: %s", line)

			break
		}

		stationName := lines[0]
		temperature := lines[1]

		_ = stationName
		_ = temperature
	}
}
