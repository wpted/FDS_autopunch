package main

import (
	"fdsAutoPunch/user"
	"fmt"
	"os"
	"time"
)

func main() {
	userName := os.Getenv("USER")
	userPassword := os.Getenv("USERPWD")
	currentUser := user.NewUser(userName, userPassword)

	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())

	if now.Before(noon) {
		if currentUser.Punch("2", "S") {
			fmt.Printf("Clock-in successfully, %s", currentUser.Account)
		}

	} else {
		if currentUser.Punch("2", "E") {
			fmt.Printf("Clock-out successfully, %s", currentUser.Account)
		}
	}
}
