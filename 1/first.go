package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main1() {
	var fuelTotal int

	file, err := os.Open("first.in")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel := (mass / 3) - 2

		fuelTotal += fuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuelTotal)
}
