package src

import (
	"bytes"
	"fmt"
	"os/exec"
)

func GetRepositories() {
	var out bytes.Buffer
	repoGitUrls := [...]string{"RioruCorporation/Rioru"}
	for _, repo := range repoGitUrls {
		exec.Command("cd", "repositories/").Run()
		cmd := exec.Command("git", "clone", "https://github.com/"+repo)
		cmd.Stdout = &out
		run := cmd.Run()
		if run != nil {
			fmt.Println(run)
		} else {
			fmt.Println("The repository \"" + repo + "\" downloaded successfully.")
		}
	}
}
