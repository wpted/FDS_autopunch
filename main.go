package main

import (
	"fdsAutoPunch/user"
	"fmt"
	"os"
	"time"
)

func main() {
	holiday2023 := map[string][]int{
		"April":     {3, 4, 5},
		"May":       {1},
		"June":      {16, 17},
		"September": {29},
		"October":   {9, 10},
		"December":  {20, 21, 22, 25, 26, 27, 28, 29},
	}

	userName := os.Getenv("USER")
	userPassword := os.Getenv("USERPWD")
	currentUser := user.NewUser(userName, userPassword)

	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	nowMonth := now.Month().String()
	nowDate := now.Day()

	// months that is special
	if dates, ok := holiday2023[nowMonth]; ok {
		// date that is special
		for _, date := range dates {
			if nowDate == date {
				fmt.Printf("You're on holiday baby")
				os.Exit(0)
			}
		}
	} else {
		if now.Before(noon) {
			if currentUser.Punch("S") {
				fmt.Printf("Clock-in successfully, %s", currentUser.Account)
			}

		} else {
			if currentUser.Punch("E") {
				fmt.Printf("Clock-out successfully, %s", currentUser.Account)
			}
		}
	}

}
