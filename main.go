package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dayOnePartOne()
	dayOnePartTwo()
}

func dayOnePartOne() {
	dial := 50
	zeroCount := 0

	f, err := os.Open("data/data_1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		step, err := strconv.Atoi(string(line[1:]))

		if nil != err {
			panic(err)
		}

		step = step % 100

		if "L" == direction {
			dial -= step
		}
		if "R" == direction {
			dial += step
		}

		if dial < 0 {
			dial = 100 + dial
		} else if dial > 99 {
			dial = dial % 100
		}

		if 0 == dial {
			zeroCount++
		}
	}

	fmt.Printf("Password: %d\n", zeroCount)
}

func dayOnePartTwo() {
	dial := 50
	zeroCount := 0

	f, err := os.Open("data/data_1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		direction := string(line[0])
		step, err := strconv.Atoi(string(line[1:]))

		if nil != err {
			panic(err)
		}

		beginZero := dial == 0

		stepZeroCounter := step / 100
		step = step % 100
		outFlag := false

		if "L" == direction {
			dial -= step
		}
		if "R" == direction {
			dial += step
		}

		if dial < 0 {
			dial = 100 + dial
			outFlag = true
		} else if dial > 99 {
			dial = dial % 100
			outFlag = true
		}

		if outFlag && 0 != dial && !beginZero {
			stepZeroCounter++
		}

		if 0 == dial {
			stepZeroCounter++
		}

		fmt.Printf("stepZero: %d / dial: %d\n", stepZeroCounter, dial)

		zeroCount += stepZeroCounter
	}

	fmt.Printf("Password: %d\n", zeroCount)
}
