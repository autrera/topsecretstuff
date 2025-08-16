package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go DD/MM/YY")
		os.Exit(1)
	}

	startDateStr := os.Args[1]
	startDate, err := parseDate(startDateStr)
	if err != nil {
		fmt.Printf("Error parsing date: %v\n", err)
		os.Exit(1)
	}

	today := time.Now()

	// Initialize git repository if not already done
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		runCommand("git", "init")
	}

	// Generate commits for each day from start date to today (excluding Sundays)
	for d := startDate; d.Before(today) || d.Equal(today); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Sunday || d.Weekday() == time.Saturday {
			continue // Skip Weekends
		}

		offDay := rand.Intn(20)
		numCommits := rand.Intn(3) + 1 // Random number between 1-3
		if offDay == 0 {
			numCommits = 0
		}

		for i := 0; i < numCommits; i++ {
			commitDate := d.Add(time.Duration(rand.Intn(24)) * time.Hour)
			commitDate = commitDate.Add(time.Duration(rand.Intn(60)) * time.Minute)

			commitMessage := fmt.Sprintf("Commit on %s", d.Format("2006-01-02"))

			runCommand("git", "commit", "--allow-empty", "--date", commitDate.Format(time.RFC3339), "-m", commitMessage)
		}
	}

	fmt.Println("Commits generated successfully!")
}

func parseDate(dateStr string) (time.Time, error) {
	parts := strings.Split(dateStr, "/")
	if len(parts) != 3 {
		return time.Time{}, fmt.Errorf("invalid date format, expected DD/MM/YY")
	}

	day := parts[0]
	month := parts[1]
	year := parts[2]

	// Handle 2-digit year (add 2000 if < 100)
	if len(year) == 2 {
		year = "20" + year
	}

	return time.Parse("02/01/2006", fmt.Sprintf("%s/%s/%s", day, month, year))
}

func runCommand(name string, args ...string) {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running %s: %v\n", name, err)
	}
}
