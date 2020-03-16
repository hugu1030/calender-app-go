package main

import (
	"fmt"

	"./database"
)

func main() {
	database.New()
	schedules, _ := database.DB.FindAllSchedules()
	fmt.Println(schedules)
}