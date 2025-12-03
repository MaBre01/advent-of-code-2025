package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//dayOnePartOne()
	//dayOnePartTwo()
	//dayTwoPartOne()
	//dayTwoPartTwo()
	//dayThreePartOne()
	dayThreePartTwo()
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

func dayTwoPartOne() {
	content, err := os.ReadFile("data/data_2.txt")
	if nil != err {
		panic(err)
	}

	stringContent := string(content)
	ranges := strings.SplitSeq(stringContent, ",")
	counter := 0

	for ran := range ranges {
		allIds := strings.Split(ran, "-")

		if len(allIds) != 2 {
			panic("The range must contain one dash")
		}

		firstId, err := strconv.Atoi(allIds[0])
		if nil != err {
			panic(err)
		}

		lastId, err := strconv.Atoi(allIds[1])
		if nil != err {
			panic(err)
		}

		for i := firstId; i <= lastId; i++ {
			current := strconv.Itoa(i)
			if len(current)%2 == 1 {
				continue
			}

			halfLen := len(current) / 2
			left := current[0:halfLen]
			right := current[halfLen:]

			if strings.Compare(left, right) == 0 {
				counter += i
			}
		}
	}

	fmt.Printf("Count: %d\n", counter)
}

func dayTwoPartTwo() {
	content, err := os.ReadFile("data/data_2.txt")
	if nil != err {
		panic(err)
	}

	stringContent := string(content)
	ranges := strings.SplitSeq(stringContent, ",")
	counter := 0

	for ran := range ranges {
		allIds := strings.Split(ran, "-")

		if len(allIds) != 2 {
			panic("The range must contain one dash")
		}

		firstId, err := strconv.Atoi(allIds[0])
		if nil != err {
			panic(err)
		}

		lastId, err := strconv.Atoi(allIds[1])
		if nil != err {
			panic(err)
		}

	out:
		for i := firstId; i <= lastId; i++ {
			current := strconv.Itoa(i)
			stringLen := len(current)
			divisor := 2

			for divisor <= stringLen {
				if stringLen%divisor != 0 {
					divisor++
					continue
				}

				invalid := true
				length := stringLen / divisor
				comparedString := current[0:length]
				for j := length; j < stringLen; j += length {
					if strings.Compare(comparedString, current[j:j+length]) != 0 {
						invalid = false
					}
				}

				if invalid {
					counter += i
					continue out
				}

				divisor++
			}
		}
	}

	fmt.Printf("Count: %d\n", counter)
}

func dayThreePartOne() {
	f, err := os.Open("data/data_3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	joltage := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		leftPosition := 0
		left, err := strconv.Atoi(string(line[leftPosition]))
		if nil != err {
			panic(err)
		}

		for i := leftPosition; i < lineLen-1; i++ {
			newLeft, err := strconv.Atoi(string(line[i]))
			if nil != err {
				panic(err)
			}

			if newLeft > left {
				left = newLeft
				leftPosition = i
			}
		}

		rightPosition := leftPosition + 1
		right, err := strconv.Atoi(string(line[rightPosition]))
		if nil != err {
			panic(err)
		}

		for i := rightPosition; i < lineLen; i++ {
			newRight, err := strconv.Atoi(string(line[i]))
			if nil != err {
				panic(err)
			}

			if newRight > right {
				right = newRight
				rightPosition = i
			}
		}

		joltage += left*10 + right
	}

	fmt.Printf("total output joltage: %d\n", joltage)
}

func dayThreePartTwo() {
	f, err := os.Open("data/data_3.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	totalJoltage := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		batteryStack := make([]string, 0)
		numberPosition := 0

		for len(batteryStack) < 12 {
			number, err := strconv.Atoi(string(line[numberPosition]))
			if nil != err {
				panic(err)
			}

			for i := numberPosition; i <= lineLen-12+len(batteryStack); i++ {
				newNumber, err := strconv.Atoi(string(line[i]))
				if nil != err {
					panic(err)
				}

				if newNumber > number {
					number = newNumber
					numberPosition = i
				}
			}

			numberString := strconv.Itoa(number)
			batteryStack = append(batteryStack, numberString)
			numberPosition++
		}

		stringJoltage := strings.Join(batteryStack, "")
		joltage, err := strconv.Atoi(stringJoltage)
		if nil != err {
			panic(err)
		}

		totalJoltage += joltage
	}

	fmt.Printf("total output joltage: %d\n", totalJoltage)
}
