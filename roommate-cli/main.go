package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/deckarep/gosx-notifier"
)

func main() {
	repl()
}

//func service() {
//
//}

func repl() {
	var line string
	var err error

	fmt.Printf("Computer Roommate Terminal\n" +
		"==========================\n" +
		"Type \"help\" for a list of commands.\n\n")

	for {
		fmt.Printf("> ")

		in := bufio.NewReader(os.Stdin)
		if line, err = in.ReadString('\n'); err != nil {
			log.Fatal(err)
		}

		args := strings.Split(strings.TrimSpace(line), " ")

		switch {
		case args[0] == "notify":
			n := &Notification{
				Body:     "Can you pick some up?",
				Title:    "Hey buddy!",
				Subtitle: "We're out of your milk.",
				Image:    "./img/angry.ico",
			}

			err := n.notify()
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("Notification sent.")
		case args[0] == "wallpaper" && len(args) > 1:
			changeWallpaper(args[1])

			fmt.Println("Wallpaper updated.")
		case args[0] == "quicklook" && len(args) > 1:
			_, err := quickLook(args[1])
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("Quicklook popped.")
		case args[0] == "startaudio" && len(args) > 1:
			go startAudio(args[1])

			fmt.Println("Audio playing.")
		case args[0] == "stopaudio":
			_, err := stopAudio()
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("Audio stopped.")
		case args[0] == "file" && len(args) > 1:
			_, err := openFile(args[1])
			if err != nil {
				fmt.Println(err.Error())
			}

			fmt.Println("That thing opened.")
		case args[0] == "help":
			fmt.Println("Valid commands:\n" +
				"notify\n" +
				"wallpaper <absolute path to image>\n" +
				"quicklook <absolute path to image>\n" +
				"file <absolute path to file>\n")
		case args[0] == "quit":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid command. Try help for a command list.")
		}
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

type Notification struct {
	Body     string
	Title    string
	Subtitle string
	Image    string
	Icon     string
}

func (this *Notification) notify() error {
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

	return note.Push()
}

func stopAudio() (string, error) {
	cmd := exec.Command("killall", "afplay")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}

func startAudio(filePath string) (string, error) {
	cmd := exec.Command("afplay", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}

func openFile(filePath string) (string, error) {
	cmd := exec.Command("open", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}

func openApp(appName string) (string, error) {
	cmd := exec.Command("open", "-a", appName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}

func quickLook(imgPath string) (string, error) {
	cmd := exec.Command("qlmanage", "-p", imgPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
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
		return "", err
	}

	list := out.String()
	return list, nil
}
