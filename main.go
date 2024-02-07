package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	git "github.com/go-git/go-git/v5"
)

var states = map[string]map[string]string{}

type repo struct {
	Name      string     `json:"Name"`
	TagStates []TagState `json:"TagStates"`
}

type TagState struct {
	Tag   string `json:"Tag"`
	State string `json:"State"`
}

func main() {
	repos := []repo{
		{Name: "neilnaveen/Test2"},
		{Name: "gittuf/gittuf"},
		{Name: "ossf/scorecard"},
	}

	for _, r := range repos {
		repoName := r.Name
		repoPath := filepath.Join("/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/", repoName)
		_, err := os.Stat(filepath.Join(repoPath, ".git"))
		if os.IsNotExist(err) {
			// If not, clone the repository
			_, err := git.PlainClone(repoPath, false, &git.CloneOptions{
				URL: "https://github.com/" + repoName + ".git",
			})

			if err != nil {
				log.Fatalf("Could not clone: %v", err)
			}

			cmd := exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/gittufinit.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])

			var stdout, stderr bytes.Buffer
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", stderr.String())
			}

		} else if err != nil {
			log.Fatalf("Could not access repository: %v", err)
		} else {
			// If the directory exists, open the repository
			_, err := git.PlainOpen(repoPath)
			if err != nil {
				log.Fatalf("Could not open existing repository: %v", err)
			}
		}

		// Pull the latest changes from the origin remote and merge into the current branch

		cmd := exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/pullUpstream.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])
		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", stderr.String())
		}

		// creating RSL entry's for the pulled in changes, we are currently only tracking tags
		// this first time we are creating RSL entries, so that if new tags have been added,
		// when gittuf finds the state for that tag commit, it will first find this RSl entry,
		// where the new tags are valid

		cmd = exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/recordRSL.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", stderr.String())
		}
		// Create rules for each of the tags, so that they are not allowed to be changed

		cmd = exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/createTagRules.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", stderr.String())
		}

		//creating RSL entry's for the pulled in changes, we are currently only tracking tags

		cmd = exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/recordRSL.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", stderr.String())
		}

		// verifying entries

		cmd = exec.Command("sh", "/Users/neilnaveen/go/src/github.com/neilnaveen/gittufchecker/scripts/verifyTagRefs.sh", strings.Split(repoName, "/")[0], strings.Split(repoName, "/")[1])

		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("Error for repo", repoName, ":", stderr.String())
		} else {
			lines := strings.Split(stdout.String(), "\n")
			for lineInd := 2; lineInd < len(lines)-1; lineInd++ {

				line := lines[lineInd]
				t, state := strings.Split(line, ": ")[0], strings.Split(line, ": ")[1]
				state = strings.Trim(state, "\n")
				if states[repoName] == nil {
					states[repoName] = map[string]string{}
				}
				states[repoName][t] = state
			}

		}
	}

	for repoIndex, r := range repos {
		for tag, state := range states[r.Name] {
			repos[repoIndex].TagStates = append(repos[repoIndex].TagStates, TagState{tag, state})
		}
	}
	repoBytes, err := json.Marshal(repos)

	if err != nil {
		fmt.Println("Error Marshaling Repo Output")
	}
	fmt.Println(string(repoBytes))
}
