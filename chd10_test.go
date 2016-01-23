package main

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"

	"github.com/deckarep/gosx-notifier"
)

func main() {
	for {
		cmd := exec.Command("ps", "-ax")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		list := out.String()
		if matched, err := regexp.MatchString("MacOS/iTunes\\b", list); err == nil && matched {
			notify()
		}

		if matched, err := regexp.MatchString("MacOS/Safari\\b", list); err == nil && matched {
			notify()
		}
	}
}

func notify() {
	note := gosxnotifier.NewNotification("Woah! I'm in here!")
	note.Title = "Give me a second."
	note.Sender = "com.apple.Safari"
	note.Sound = gosxnotifier.Basso
	note.Push()
}

// Change Desktop Wallpaper
// osascript -e 'tell application "Finder" to set desktop picture to POSIX file "/path/to/picture.jpg"'
