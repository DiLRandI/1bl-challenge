package app

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const minLineLength = 2

type WeatherData struct {
	StationName string
	Temperature int64
}

func Execute(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	stationTemperatures := map[string][]string{}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("failed to read line: %v", err)
		}

		lines := strings.Split(line, ";")

		if len(lines) < minLineLength {
			continue
		}

		stationName := lines[0]
		temperature := lines[1]

		stationTemperatures[stationName] = append(stationTemperatures[stationName], temperature)
	}

	return nil
}
