package main

import (
	"os"
	"log"
	"bufio"

	"github.com/golang-collections/collections/set"
	"math"
	"strings"
	"strconv"
	"fmt"
)

type Point struct {
	x int
	y int
}

func getWireLocations(wire string) (*set.Set, error){
	wirePath := getWirePath(wire)
	wireLocation := set.New()

	currPoint := Point{0, 0}
	for _, nextPath := range wirePath {
		direction, steps, err := readDirection(nextPath)
		if err != nil {
			return nil, err
		}
		//for i := 0; i < steps; i++ {
		//	switch direction {
		//	case "U":
		//		currPoint.y += 1
		//	case "D":
		//		currPoint.y -= 1
		//	case "L":
		//		currPoint.x -= 1
		//	case "R":
		//		currPoint.x += 1
		//	}
		//	wireLocation.Insert(currPoint)
		//}

		switch direction {
		case "U":
			for i := 0; i < steps; i++ {
				currPoint.y += 1
				wireLocation.Insert(currPoint)
			}
		case "D":
			for i := 0; i < steps; i++ {
				currPoint.y -= 1
				wireLocation.Insert(currPoint)
			}
		case "L":
			for i := 0; i < steps; i++ {
				currPoint.x -= 1
				wireLocation.Insert(currPoint)
			}
		case "R":
			for i := 0; i < steps; i++ {
				currPoint.x += 1
				wireLocation.Insert(currPoint)
			}
		}

	}
	return wireLocation, nil
}

func readDirection(path string) (string, int, error) {
	direction := string(path[0])
	strNumPoints := path[1:]
	numPoints, err := strconv.Atoi(strNumPoints)
	if err != nil {
		log.Printf("unexpected value when trying to read directions, expected integers but was %v", strNumPoints)
		return "", 0, err
	}
	return direction, numPoints, nil
}

func getWirePath(wireIn string) []string {
	wireIn = strings.TrimSpace(wireIn)
	return strings.Split(wireIn, ",")
}

func getAbsVal(val int) int {
	if val < 0 {
		return -1 * val
	}
	return val
}

func findClosestIntersectionDist(intersections *set.Set) int {
	closest := math.MaxInt64
	intersections.Do(func(intersection interface{}) {
		point := intersection.(Point)
		currDist := getAbsVal(point.x) + getAbsVal(point.y)
		if currDist < closest {
			closest = currDist
		}
	})
	return closest
}

func readInput(path string) ([]string, error){
	file, err := os.Open(path)
	if err != nil {
		log.Printf("error trying to open file at %v", path)
		return nil, err
	}
	defer file.Close()

	var wires []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wires = append(wires, scanner.Text())
	}

	return wires, nil
}

func main() {
	path := os.Args[1]

	wires, err := readInput(path)
	if err != nil {
		log.Fatal(err)
	}

	wireLocs1, err := getWireLocations(wires[0])
	if err != nil {
		log.Fatal(err)
	}
	wireLocs2, err := getWireLocations(wires[1])
	if err != nil {
		log.Fatal(err)
	}

	intersections := wireLocs1.Intersection(wireLocs2)
	dist := findClosestIntersectionDist(intersections)

	fmt.Println("closest intersection to center is: ", dist)
}