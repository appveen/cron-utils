package main

import (
	utils "./utils"
	"fmt"
)

func main() {
	timestamp := "2019-12-18T10:08:44.630Z"
	cronRegEx := "* 17-20 * * 1-5"
	timebounds := make([]utils.Timebound, 0)
	timebounds = append(timebounds, utils.Timebound{
		From: "17:17",
		To:   "22:45",
	})
	if utils.CheckForTimeBound(cronRegEx, timestamp, timebounds) == true {
		fmt.Println("Valid Time")
	} else {
		fmt.Println("In-valid Time")
	}
}
