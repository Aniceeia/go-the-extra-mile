package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Visit struct {
	Specialization string
	VisitDate      time.Time
}

type UserNotFoundError struct {
	message string
}

func (e *UserNotFoundError) Error() string {
	return e.message
}

func Save(note *map[string][]Visit, name, specialization, inputDate string, lastVisitMemory map[string]time.Time) error {
	date, err := time.Parse("2006-01-02", inputDate)
	if err != nil {
		return errors.New("Invalide date format")
	}

	cacheKey := name + " " + specialization
	lastVisitMemory[cacheKey] = date

	visit := Visit{Specialization: specialization, VisitDate: date}
	(*note)[name] = append((*note)[name], visit)

	return nil
}

func GetHistory(name string, note map[string][]Visit) ([]Visit, error) {
	visits, names := note[name]
	if !names {
		return nil, &UserNotFoundError{message: "user not found"}
	}

	return visits, nil
}

func GetLastVisit(name, specialization string, note map[string][]Visit, lastVisitMemory map[string]time.Time) (time.Time, error) {
	var lastDate time.Time
	visits, names := note[name]
	cacheKey := name + " " + specialization
	if visit, exists := lastVisitMemory[cacheKey]; exists {
		return visit, nil
	}
	if !names {
		return time.Time{}, &UserNotFoundError{message: "user not found"}
	}

	for _, val := range visits {
		if val.Specialization == specialization && val.VisitDate.After(lastDate) {
			lastDate = val.VisitDate
		}
	}

	return lastDate, nil
}

func readJournal() (string, string, string) {
	var str, buff string
	for range 5 {
		fmt.Scanf("%s", &buff)
		str += buff + " "
	}
	splittedString := strings.Split(str, " ")
	name := splittedString[0] + " " + splittedString[1] + " " + splittedString[2]
	doctor := splittedString[3]
	date := splittedString[4]

	return name, doctor, date
}

func journaling(operation string, journal map[string][]Visit, cache map[string]time.Time) {
	for {
		fmt.Scanf("%s", &operation)
		switch operation {
		case "Save":
			name, doctor, date := readJournal()
			Save(&journal, name, doctor, date, cache)
		case "GetHistory":
			var buff, str string
			for range 3 {
				fmt.Scanf("%s", &buff)
				str += buff + " "
			}
			splittedStr := strings.Split(str, " ")
			historyName := splittedStr[0] + " " + splittedStr[1] + " " + splittedStr[2]
			visits, err := GetHistory(historyName, journal)
			if (err != nil || visits == nil){
				fmt.Println(&UserNotFoundError{message: "user not found"})
				return
			}
			for _, v := range visits {
				fmt.Println(v.Specialization, v.VisitDate)
			}
		case "GetLastVisit":
			var buff, str string
			for range 4 {
				fmt.Scanf("%s", &buff)
				str += buff + " "
			}
			splittedStr := strings.Split(str, " ")
			lastVisitName := splittedStr[0] + " " + splittedStr[1] + " " + splittedStr[2]
			lastVisitSpecialist := splittedStr[3]
			date, err := GetLastVisit(lastVisitName, lastVisitSpecialist, journal, cache)
			if (err != nil || date == time.Time{}) {
				fmt.Println(&UserNotFoundError{message: "user not found"})
				return
			}
			fmt.Println(date)
		}
	}
}

func main() {
	journal := make(map[string][]Visit)
	cache := make(map[string]time.Time)
	var operation string
	journaling(operation, journal, cache)
}
