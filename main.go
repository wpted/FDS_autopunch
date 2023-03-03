package main

import (
	"fdsAutoPunch/user"
	"fmt"
	"os"
	"time"
)

func main() {
	userName := os.Getenv("USERNAME")
	userPassword := os.Getenv("USERPWD")
	currentUser := user.NewUser(userName, userPassword)

	now := time.Now()
	noon := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())

	if now.Before(noon) {
		fmt.Printf("It's before noon, %s", currentUser.Account)
		currentUser.Punch("2", "S")

	} else {
		fmt.Printf("It's after noon, %s", currentUser.Account)
		currentUser.Punch("2", "E")
	}
}
