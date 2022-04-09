package service

import (
	"ayush-deliveroo/internal/utils"
	"fmt"
	"strings"
)

func Cron(args []string) {
	cronCommand := args[0]
	cronSplit := strings.Split(cronCommand, " ")

	if len(cronSplit) != 6 {
		fmt.Println("Error")
	} else {
		min := utils.HandleCronFormat(cronSplit[0], "minute")
		hour := utils.HandleCronFormat(cronSplit[1], "hour")
		dayOfMonth := utils.HandleCronFormat(cronSplit[2], "dayOfMonth")
		month := utils.HandleCronFormat(cronSplit[3], "month")
		dayOfWeek := utils.HandleCronFormat(cronSplit[4], "dayOfWeek")
		command := cronSplit[5]

		utils.Print("minute        ", min)
		utils.Print("hour          ", hour)
		utils.Print("day of month  ", dayOfMonth)
		utils.Print("month         ", month)
		utils.Print("day of week   ", dayOfWeek)
		fmt.Println("command      ", command)
	}
}
