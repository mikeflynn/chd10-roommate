package main

import (
	"fmt"

	"github.com/deckarep/gosx-notifier"
)

var EventList map[string]*Event = map[string]*Event{
	"wallpaper": &Event{
		Description:    "Changes the wallpaper on the desktop.",
		ArgDescription: "<absolute path to image>",
		Attributes: AttributeList{
			Clean:  50,
			Angry:  50,
			Creepy: 50,
		},
		Options: {},
		Fn: func(...string) string {
			cmd := fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", imgPath)
			actionScript(cmd)
		},
	},
	"notify": &Event{
		ArgDescription: "",
	},
	"quicklook": &Event{
		ArgDescription: "<absolute path to image>",
	},
	"movefile": &Event{
		ArgDescription: "<absolute path to source> <absolute path to desitination>",
	},
	"openfile": &Event{
		ArgDescription: "<absolute path to file>",
	},
	"makedir": &Event{
		ArgDescription: "<absolute path to new directory>",
	},
	"makefile": &Event{
		ArgDescription: "<absolute path to new file> <number of chars to fill it with>",
	},
	"openapp": &Event{
		ArgDescription: "<app name> <background flag>",
	},
	"closeapp": &Event{
		ArgDescription: "<app name>",
	},
	"brightness": &Event{
		ArgDescription: "<brightness level 0 - 1; ex: 0.3>",
	},
	"alert": &Event{
		ArgDescription: "<body> <title> <icon path> <button_1_text> <button_2_text>",
	},
	"volume": &Event{
		ArgDescription: "<volume % 0 - 100>",
	},
	"startaudio": &Event{
		ArgDescription: "<absolute path to audio>",
	},
	"stopaudio": &Event{
		ArgDescription: "",
	},
	"commands": &Event{
		ArgDescription: "",
	},
}

type Event struct {
	Description     string                 // Name of event for logging.
	ArgDescription  string                 // A description of what arguments it takes.
	Attributes      AttributeList          // Descriptors
	Options         map[string]string      // Various options.
	Fn              func(...string) string // Run method
	FollowedBy      []*Event               // Links to events that may come directly after this one.
	LastOccured     uint64                 // Time stamp of last occurrence.
	DownTime        uint64                 // Min number of seconds between occurrences.
	TotalOccurances int                    // Total number of occurrences in this run.
}

func (this *Event) Run(args ...string) {
	fmt.Println(this.Fn(args...))
}

// Utility Functions

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

func storedActionScript(scriptName string, params ...string) (string, error) {
	var data []byte
	var err error

	// Pull from asset store
	if data, err = Asset("scripts/" + scriptName); err != nil {
		return "", err
	}

	if err = ioutil.WriteFile("/tmp/"+scriptName, []byte(data), 0644); err != nil {
		return "", err
	}

	if err = os.Chmod("/tmp/"+scriptName, 0777); err != nil {
		return "", err
	}

	cmd := exec.Command("/tmp/"+scriptName, params...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
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

func termCommand(args ...string) (string, error) {
	cmd := exec.Command(args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}
