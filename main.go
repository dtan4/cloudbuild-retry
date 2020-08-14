package main

import (
	"fmt"
	"os"
)

const (
	exitOK    = 0
	exitError = 1
)

func run(args []string) int {
	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "Cloud Build project and build ID is required")
		return exitError
	}
	projectID, buildID := args[1], args[2]

	fmt.Printf("project: %s\n", projectID)
	fmt.Printf("build: %s\n", buildID)

	return exitOK
}

func main() {
	os.Exit(run(os.Args))
}
