package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"strconv"
	"errors"
)

func sliceToMap(slice1 []int) map[int]bool {
	mappedSlice := make(map[int]bool)
	for _, v := range slice1 {
		mappedSlice[v] = true
	}

	return mappedSlice
}

func findInptersection(slice2 []int, mappedSlice map[int]bool) []int {
	var result []int
	for _, v := range slice2 {
		if mappedSlice[v] {
			result = append(result, v)
			mappedSlice[v] = false
		}
	}

	return result
}

func reader() ([]int, error) {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	numbersStr := strings.Fields(str)
	if str == "" {
		return nil, errors.New("Invalid input")
	}
	var slicedStr []int
	for _, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, errors.New("Invalid input")
		}
		slicedStr = append(slicedStr, num)
	}

	return slicedStr, nil
}

func main() {
	slice1, err := reader()
	if err != nil {
		fmt.Println(err)
		return
	}
	slice2, err := reader()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := findInptersection(slice1, sliceToMap(slice2))
	if len(result) == 0 {
		err := errors.New("Empty intersection")
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}