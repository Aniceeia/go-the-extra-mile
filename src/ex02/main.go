package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type SlicedMap struct {
	Word  string
	Count int
}

func wordCount(str string) map[string]int {
	mapWord := make(map[string]int)

	strTrimmed := strings.Fields(str)
	for _, val := range strTrimmed {
		mapWord[val] += 1
	}

	return mapWord
}

func mapToSlice(mapWord map[string]int) []SlicedMap {
	var wordCount []SlicedMap
	for i, val := range mapWord {
		wordCount = append(wordCount, SlicedMap{i, val})
	}

	return wordCount
}

func sortedSlice(mapToSort map[string]int) []SlicedMap {
	slicedMap := mapToSlice(mapToSort)
	sort.Slice(slicedMap, func(i, j int) bool {
		if slicedMap[i].Count == slicedMap[j].Count {
			return slicedMap[i].Word < slicedMap[j].Word
		}
		return slicedMap[i].Count > slicedMap[j].Count
	})

	return slicedMap
}

func handleInput(str string, ammount int) ([]SlicedMap, error) {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return nil, errors.New("ERROR: empty string")
	}

	if ammount <= 0 {
		return nil, errors.New("ERROR: ammount of words less than 0")
	}

	mapOfStr := wordCount(str)
	if ammount > len(mapOfStr) {
		ammount = len(mapOfStr)
	}

	return sortedSlice(mapOfStr)[0:ammount], nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	var ammount int
	_, err := fmt.Scan(&ammount)
	if err != nil {
		fmt.Println("ERROR: invalid input for K")
		return
	}

	result, err := handleInput(str, ammount)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, val := range result {
		if i < ammount-1 {
			fmt.Printf("%s ", val.Word)
		} else {
			fmt.Printf("%s", val.Word)
		}
	}
}
