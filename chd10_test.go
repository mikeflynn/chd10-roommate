package main

import (
	"bytes"
	"log"
	"os/exec"
	"regexp"

	"github.com/deckarep/gosx-notifier"
)

func main() {

}

func notify() {
	note := gosxnotifier.NewNotification("Woah! I'm in here!")
	note.Title = "Give me a second."
	note.Sender = "com.apple.Safari"
	note.Sound = gosxnotifier.Basso
	note.Push()
}

func watchPs() {
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

func changeWallpaper(imgPath string) {
	return exec.Command("tell application \"Finder\" to set desktop picture to POSIX file \"/path/to/picture.jpg\"")
}

func actionScript(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", "'"+command+"'")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}

	list := out.String()
	return list, nil
}

// Change Desktop Wallpaper
// osascript -e 'tell application "Finder" to set desktop picture to POSIX file "/path/to/picture.jpg"'

// Quick Look
// qlmanage -p ~/Desktop/spider-meme-20.jpg
