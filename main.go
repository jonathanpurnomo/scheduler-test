package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		timeStr := time.Now().Format("01-02-2006 15:04")
		key, message := os.Args[1], strings.Join(os.Args[2:], " ")

		if strings.EqualFold(key, "ping") {
			prompResult := fmt.Sprintf("%s %s %s", "pong triggered using", message, timeStr)
			fmt.Println(prompResult)
		} else {
			fmt.Println("Invalid input")
		}
	} else {
		fmt.Println("No input provided")
	}
}
