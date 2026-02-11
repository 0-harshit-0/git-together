package main

import (
	"log"
	"os/exec"
	"runtime"

	"github.com/gen2brain/dlgs"
)

const repoURL = "https://github.com/0-harshit-0/git-together?tab=readme-ov-file#the-3-way-handshake"

func main() {
	answer, err := dlgs.Question(
		"Important Question",
		"Will you be my valentine?",
		true,
	)

	if err != nil {
		log.Fatal(err)
	}

	if answer {
		openBrowser(repoURL)
		return
	}

	secondBox()
}

func secondBox() {
	for {
		answer, err := dlgs.Question(
			"Are you sure?",
			"That did not seem right.\nWill you be my valentine?",
			true,
		)

		if err != nil {
			log.Fatal(err)
		}

		if answer {
			openBrowser(repoURL)
			return
		}

		// If NO is clicked, loop again
		// If the user closes the dialog, program exits naturally
	}
}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
}
