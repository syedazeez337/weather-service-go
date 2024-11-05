package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"weatherservice/weather"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the weather Tracker!")
	fmt.Println("Type 'quit' to exit the application.\n")

	for {
		fmt.Print("Enter city name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read input: %v", err)
		}

		city := strings.TrimSpace(input)
		if city == "quit" {
			fmt.Println("Exiting Weather Tracker...")
			break
		}

		data, err := weather.GetWeather(city)
		if err != nil {
			fmt.Printf("Error fetching weather: %v\n", err)
			continue
		}

		weather.DisplayWeather(data)
	}
}
