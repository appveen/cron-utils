package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//Timebound Structure of timebounds array
type Timebound struct {
	From string
	To   string
}

//CheckForTimeBound Will return true if the timestamp is within Timebound
func CheckForTimeBound(cronRegEx string, timestamp string, timebounds []Timebound) bool {
	t := time.Parse(time.RFC3339, timestamp)
	minute := t.Local().Minute()
	hour := t.Local().Hour()
	date := t.Local().Day()
	month := int(t.Local().Month())
	weekDay := int(t.Local().Weekday())
	var flag bool = true
	validMinutes := getValidValues(cronRegEx, 0, 0, 59)
	validHours := getValidValues(cronRegEx, 1, 0, 23)
	validDate := getValidValues(cronRegEx, 2, 1, 31)
	validMonth := getValidValues(cronRegEx, 3, 1, 12)
	validWeekDay := getValidValues(cronRegEx, 4, 0, 7)
	if timebounds != nil {
		currTime := (hour * 100) + minute
		flag = false
		for _, item := range timebounds {
			lowerLimit, _ := strconv.Atoi(strings.ReplaceAll(item.From, ":", ""))
			upperLimit, _ := strconv.Atoi(strings.ReplaceAll(item.To, ":", ""))
			fmt.Println(lowerLimit)
			fmt.Println(currTime)
			fmt.Println(upperLimit)
			if currTime >= lowerLimit && currTime <= upperLimit {
				flag = true
			}
		}
	} else {
		if !contains(validMinutes, minute) {
			flag = false
		}
		if !contains(validHours, hour) {
			flag = false
		}
	}
	if !contains(validDate, date) {
		flag = false
	}
	if !contains(validMonth, month) {
		flag = false
	}
	if !contains(validWeekDay, weekDay) {
		flag = false
	}
	return flag
}

func getValidValues(cronRegEx string, crontab int, minValue int, maxValue int) []int {
	segments := strings.Split(cronRegEx, " ")
	validHours := make([]int, 0)
	if segments[crontab] != "*" {
		hours := strings.Split(segments[crontab], ",")
		for _, item := range hours {
			segments2 := strings.Split(item, "-")
			num1, _ := strconv.Atoi(segments2[0])
			var num2 int = 0
			if len(segments2) > 1 {
				num2, _ = strconv.Atoi(segments2[1])
				len := num2 - num1
				for index := 0; index <= len; index++ {
					validHours = append(validHours, num1+index)
				}
			} else {
				validHours = append(validHours, num1)
			}
		}
	} else {
		for index := minValue; index <= maxValue; index++ {
			validHours = append(validHours, index)
		}
	}
	return validHours
}

func contains(arr []int, item int) bool {
	for _, a := range arr {
		if a == item {
			return true
		}
	}
	return false
}
