package utils

import (
	"fmt"
	"strconv"
	"strings"
)

/*
assumed month to be of 30 days for all the cases
*/
func HandleCronFormat(s string, field string) []string {
	if strings.Contains(s, ",") {
		return strings.Split(s, ",")
	} else if strings.Contains(s, "-") {
		return handleHyphen(s)
	} else if strings.Contains(s, "/") {
		return handleForwardSlash(s, field)
	} else {
		return handleRemainingCase(s, field)
	}
}

func handleHyphen(s string) []string {

	arrayBounds := strings.Split(s, "-")
	start, _ := strconv.Atoi(arrayBounds[0])
	end, _ := strconv.Atoi(arrayBounds[1])

	return Time[start : end+1]
}

func handleForwardSlash(s string, field string) []string {
	arrayBounds := strings.Split(s, "/")
	count, _ := strconv.Atoi(arrayBounds[1])

	switch field {
	case "minute":
		return forwardSlackLoop(count, 60)
	case "hour":
		return forwardSlackLoop(count, 24)
	case "dayOfMonth":
		return forwardSlackLoop(count, 30)
	case "month":
		return forwardSlackLoop(count, 12)
	case "dayOfWeek":
		return forwardSlackLoop(count, 7)
	default:
		return []string{}
	}
}

func forwardSlackLoop(count int, end int) []string {
	var timeArray = make([]string, 0)
	timeArray = append(timeArray, "0")
	for i := count; i < 60; i += count {
		timeArray = append(timeArray, Time[i])
	}
	return timeArray
}

func handleRemainingCase(s string, field string) []string {
	if s == "*" {
		switch field {
		case "minute":
			return Time[0:60]
		case "hour":
			return Time[0:24]
		case "dayOfMonth":
			return Time[1:30]
		case "month":
			return Time[1:13]
		case "dayOfWeek":
			return Time[1:8]
		default:
			return []string{}
		}

	} else {
		return []string{s}
	}

}

func Print(s string, field []string) {
	fmt.Print(s)
	for _, v := range field {
		fmt.Print(v, " ")
	}
	fmt.Println(" ")
}
