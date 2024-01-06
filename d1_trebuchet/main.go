package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	file, _ := os.Open("./input.txt")

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	for _,line := range fileLines {
		result := composeTwoDigit(line)
		sum += result
		fmt.Println(result)
	}
	timeElapsed := time.Since(start)

	fmt.Printf("Sum total of the numbers are %v ", sum)
	fmt.Println()
	fmt.Printf("Code Execution took %s", timeElapsed)
}

func composeTwoDigit(line string) int {
	reg, _ := regexp.Compile("[^0-9]+")
	
	numericStr := reg.ReplaceAllString(line, "")

	switch len(numericStr) {
	case 1 :
		result := numericStr + numericStr
		parsed, _ := strconv.Atoi(result)
		return parsed

	case 2:
		parsed, _ := strconv.Atoi(numericStr)
		return parsed

		default:
			result := string(numericStr[0]) + string(numericStr[len(numericStr)-1])
			parsed, _ := strconv.Atoi(result)
			return parsed
	}

}