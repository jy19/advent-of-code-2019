package main

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

const (
	steps  = 4
	output = 19690720
)

func processOps(opCodes []int) []int {

	opNum := 0
	for {
		cmds := getNextSequence(opNum, opCodes)
		if cmds[0] == 99 {
			break
		}
		switch cmds[0] {
		case 1:
			opCodes[cmds[3]] = opCodes[cmds[1]] + opCodes[cmds[2]]
		case 2:
			opCodes[cmds[3]] = opCodes[cmds[1]] * opCodes[cmds[2]]
		}
		opNum += 1
	}

	return opCodes
}

func findParameters(opCodes []int) (int, int) {
	noun := 0
	verb := 0

	for noun := range [100]int{} {
		for verb := range [100]int{} {
			cpyOps := make([]int, len(opCodes))
			copy(cpyOps, opCodes)
			cpyOps[1] = noun
			cpyOps[2] = verb
			resultOps := processOps(cpyOps)
			if resultOps[0] == output {
				return noun, verb
			}
		}
	}
	log.Printf("didn't find a result..")
	return noun, verb
}

func printProcess(opCodes []int) string {
	result := ""
	for _, op := range opCodes {
		if len(result) > 0 {
			result += ","
		}
		result += strconv.Itoa(op)
	}
	return result
}

func getNextSequence(opNum int, opCodes []int) []int {
	if (opNum*steps)+steps >= len(opCodes) {
		return opCodes[opNum*steps:]
	}
	return opCodes[opNum*steps:(opNum*steps)+steps]
}

func readOpArray(path string) ([]int, error) {
	byteSlice, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("error trying to read file at %v", path)
		return nil, err
	}
	cleanedString := strings.TrimSpace(string(byteSlice))
	opArrString := strings.Split(cleanedString, ",")
	opArrInt := make([]int, len(opArrString))
	for i, op := range opArrString {
		opArrInt[i], err = strconv.Atoi(op)
		if err != nil {
			log.Printf("unexpected value in input, expected integers but was %v", op)
			return nil, err
		}
	}

	return opArrInt, nil
}

func main() {
	path := os.Args[1]

	input, err := readOpArray(path)
	if err != nil {
		log.Fatal(err)
	}

	//resultOps := processOps(input)
	//result := printProcess(resultOps)
	//fmt.Println("final state of program: ", result)

	noun, verb := findParameters(input)
	fmt.Println("parameters for output: ", noun, verb)
}