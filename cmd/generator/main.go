package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

// WeatherStation represents a weather station with a name
type WeatherStation struct {
	Name string
}

func main() {
	// Seed the random number generator
	sourceFile, err := os.Open("./data/weather_stations.csv")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer sourceFile.Close()
	stationNames := []string{}
	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		txt := scanner.Text()
		spitTxt := strings.Split(txt, ";")
		if len(spitTxt) < 2 {
			log.Printf("invalid line: %v", txt)
			continue
		}

		stationNames = append(stationNames, spitTxt[0])
	}

	// Output file name
	fileName := "./data/weather_data.csv"

	// Open the output file for writing
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Close the file on exit

	// Generate 1 billion data points
	var wg sync.WaitGroup
	lintCh := make(chan string, 100000000)
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		seeder := rand.New(rand.NewSource(time.Now().UnixNano()))

		go func(ch chan<- string, seeder *rand.Rand) {
			defer wg.Done()
			for j := 0; j < 1000000; j++ {
				// Generate a random temperature between -50 and 50

				temperature := seeder.Float64()*100 - 50

				// Generate a random station name
				stationName := stationNames[seeder.Intn(len(stationNames))]

				// Write the data to the file
				ch <- fmt.Sprintf("%s;%.2f", stationName, temperature)
			}
		}(lintCh, seeder)
	}

	wg.Add(1)
	go func(ch <-chan string) {
		defer wg.Done()
		for line := range ch {
			log.Printf("writing line: %v", line)
			fmt.Fprintf(file, "%s\n", line)
		}
	}(lintCh)

	wg.Wait()
	close(lintCh)

	fmt.Printf("Data generated successfully and written to file: %s\n", fileName)
}
