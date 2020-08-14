package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/api/cloudbuild/v1"
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

	ctx := context.Background()
	cloudbuildService, err := cloudbuild.NewService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	op, err := cloudbuildService.Projects.Builds.Retry(projectID, buildID, &cloudbuild.RetryBuildRequest{}).Do()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	b, err := op.Metadata.MarshalJSON()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	var m cloudbuild.BuildOperationMetadata

	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return exitError
	}

	fmt.Printf("new build: %s\n", m.Build.Id)

	return exitOK
}

func main() {
	os.Exit(run(os.Args))
}
