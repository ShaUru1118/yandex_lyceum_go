package main

import (
	"fmt"
	"time"
)

func TimeAgo(pastTime time.Time) string {
	currentTime := time.Now()
	duration := currentTime.Sub(pastTime)
	years := duration.Hours() / 24 / 365
	months := duration.Hours() / 24 / 30
	days := duration.Hours() / 24
	hours := duration.Hours()
	minutes := duration.Minutes()
	seconds := duration.Seconds()

	if years >= 1 {
		return fmt.Sprintf("%.0f years ago", years)
	} else if months >= 1 {
		return fmt.Sprintf("%.0f months ago", months)
	} else if days >= 1 {
		return fmt.Sprintf("%.0f days ago", days)
	} else if hours >= 1 {
		return fmt.Sprintf("%.0f hours ago", hours)
	} else if minutes >= 1 {
		return fmt.Sprintf("%.0f minutes ago", minutes)
	} else {
		return fmt.Sprintf("%.0f seconds ago", seconds)
	}

}

// func main() {
// 	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
// 	result := TimeAgo(pastTime)
// 	fmt.Println(result)
// }
