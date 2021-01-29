package main

import (
	"log"
	"os"
)

func main() {
	var (
		// inputs
		inputProjects = os.Getenv("DATATUG_PROJECTS")
		// github env
		githubWorkflow   = os.Getenv("GITHUB_WORKFLOW")
		githubRepository = os.Getenv("GITHUB_REPOSITORY")
		githubCommit     = os.Getenv("GITHUB_SHA")
	)

	if inputProjects == "" {
		log.Fatal("DATATUG_PROJECTS input is required")
	}

	log.Println("Workflow: ", githubWorkflow)
	log.Println("Repository: ", githubRepository)
	log.Println("Commit: ", githubCommit)
	log.Println("DATATUG_PROJECTS: ", inputProjects)
}
