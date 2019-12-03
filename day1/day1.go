package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func calculate(modules []int) int {
	sum := 0
	for _, module := range modules {
		sum += calcFuelNeeded(module)
	}
	return sum
}

func calcFuelNeeded(mass int) int {
	return (mass/3) - 2
}

func calculateFuelForFuel(fuel int) int {
	fuels := stack.New()
	fuels.Push(fuel)
	sum := 0

	for {
		if fuels.Len() == 0 {
			break
		}

		fuel := fuels.Pop()
		fuelNeeded := calcFuelNeeded(fuel.(int))
		if fuelNeeded < 0 {
			fuelNeeded = 0
		}
		sum += fuelNeeded
		if fuelNeeded > 0 {
			fuels.Push(fuelNeeded)
		}
	}
	return sum
}

func calcFuelNeededv2(mass int) int {
	fuel := (mass/3) - 2
	if fuel < 0 {
		return 0
	}
	return fuel
}

func calculateFuelv2(modules []int) int {
	sum := 0
	masses := stack.New()
	for _, module := range modules {
		fuelForModule := calcFuelNeeded(module)
		masses.Push(fuelForModule)
		for {
			if masses.Len() == 0 {
				break
			}

			mass := masses.Pop()
			currFuelNeeded := calcFuelNeededv2(mass.(int))
			fuelForModule += currFuelNeeded
			if currFuelNeeded > 0 {
				masses.Push(currFuelNeeded)
			}
		}
		sum += fuelForModule
	}
	return sum
}

func readLinesToInt(path string) ([]int, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		log.Printf("error trying to open file at %v", path)
		return nil, err
	}
	defer inputFile.Close()

	var lines []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		intLine, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("something went wrong trying to strconv scanner text %v to int", scanner.Text())
			return nil, err
		}
		lines = append(lines, intLine)
	}

	return lines, nil
}

func main() {
	path := os.Args[1]

	modules, err := readLinesToInt(path)
	if err != nil {
		log.Fatal(err)
	}

	//sum := calculate(modules)
	//fmt.Println("calculated total: ", sum)
	//
	//totalFuelNeeded := calculateFuelForFuel(sum)
	//fmt.Println("total fuel needed: ", sum+totalFuelNeeded)

	sum := calculateFuelv2(modules)
	fmt.Println("total fuel needed: ", sum)
}
