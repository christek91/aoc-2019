package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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

		fmt.Println("Fuel:", fuel)
		fuelTotal += findFuelMass(fuel)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(fuelTotal)
}

// Find the additional fuel needed to ship the input amount of fuel to space.
func findFuelMass(fuelReq int) int {
	neededFuel := (fuelReq / 3) - 2
	if neededFuel <= 0 {
		return fuelReq
	}

	return findFuelMass(neededFuel) + fuelReq
}
