package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/deckarep/gosx-notifier"
)

func main() {
	var command string

	for {
		fmt.Printf("What should I do:")
		if _, err := fmt.Scan(&command); err != nil {
			log.Fatal(err)
		}

		if command == "notify" {
			n := &Notification{
				Body:     "Can you pick some up?",
				Title:    "Hey buddy!",
				Subtitle: "We're out of your milk.",
				Image:    "../img/angry.ico",
			}

			n.notify()
		} else if command == "wallpaper" {
			changeWallpaper("../img/wallpaper.jpg")
		} else if command == "quicklook" {
			quickLook("../img/milk.jpg")
		}
	}
}

type Notification struct {
	Body     string
	Title    string
	Subtitle string
	Image    string
	Icon     string
}

func (this *Notification) notify() {
	note := gosxnotifier.NewNotification(this.Body)
	note.Title = this.Title
	note.Subtitle = this.Subtitle
	note.Group = "com.roommate.cli.chd10"

	if this.Icon != "" {
		note.AppIcon = this.Icon
	}

	if this.Image != "" {
		note.ContentImage = this.Image
	}

	note.Sound = gosxnotifier.Basso

	err := note.Push()
	if err != nil {
		log.Println("Failed notification.")
	}
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
			//notify()
		}

		if matched, err := regexp.MatchString("MacOS/Safari\\b", list); err == nil && matched {
			//notify()
		}
	}
}

func quickLook(imgPath string) (string, error) {
	cmd := exec.Command("qlmanage", "-p", imgPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		//log.Println(err.Error())
		return "", err
	}

	list := out.String()
	return list, nil
}

func changeWallpaper(imgPath string) {
	cmd := fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", imgPath)
	actionScript(cmd)
}

func actionScript(command string) (string, error) {
	cmd := exec.Command("osascript", "-e", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		//log.Println(err.Error())
		return "", err
	}

	list := out.String()
	return list, nil
}
