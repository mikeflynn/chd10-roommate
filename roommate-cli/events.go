package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/deckarep/gosx-notifier"
)

var EventList map[string]*Event = map[string]*Event{
	"wallpaper": &Event{
		Description:    "Changes the wallpaper on the desktop.",
		ArgDescription: "<absolute path to image>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				cmd := fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", args[0])
				actionScript(cmd)
			}
			return "Wallpaper set."
		},
	},
	"notify": &Event{
		Description:    "Fires a notification to the screen.",
		ArgDescription: "<title> <body> <image>",
		Fn: func(args ...string) string {
			if len(args) < 4 {
				return "Not enough arguments."
			} else {
				n := &Notification{
					Body:  args[0],
					Title: args[1],
					Image: args[2],
				}

				err := n.notify()
				if err != nil {
					return err.Error()
				}

				return "Notification sent."
			}
		},
	},
	"quicklook": &Event{
		Description:    "Opens quicklook with a file.",
		ArgDescription: "<absolute path to image>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				cmd := exec.Command("qlmanage", "-p", args[0])
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					return err.Error()
				}

				list := out.String()
				if list != "" {
					return list
				}

				return "Quicklook popped."
			}
		},
	},
	"movefile": &Event{
		Description:    "Moves the specified file to the specified place.",
		ArgDescription: "<absolute path to source> <absolute path to desitination>",
		Fn: func(args ...string) string {
			if len(args) < 3 {
				return "Not enough arguments."
			} else {
				if err := os.Rename(args[0], args[1]); err != nil {
					return err.Error()
				} else {
					return "File moved."
				}
			}
		},
	},
	"openfile": &Event{
		Description:    "Opens the specified file.",
		ArgDescription: "<absolute path to file>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				cmd := exec.Command("open", args[0])
				var out bytes.Buffer
				cmd.Stdout = &out
				err := cmd.Run()
				if err != nil {
					return err.Error()
				}

				if list := out.String(); list != "" {
					return list
				}

				return "File opened."
			}
		},
	},
	"makedir": &Event{
		Description:    "Creates a directory.",
		ArgDescription: "<absolute path to new directory>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				if err := os.MkdirAll(args[0], 0777); err != nil {
					return err.Error()
				}

				return "Directory created."
			}
		},
	},
	"makefile": &Event{
		Description:    "Generates a file with the required number of characters.",
		ArgDescription: "<absolute path to new file> <number of chars to fill it with>",
		Fn: func(args ...string) string {
			var charCount uint64
			var filePath string
			var err error

			if len(args) == 0 {
				return "Not enough arguments."
			} else if len(args) == 1 {
				filePath = args[0]
				charCount = 100
			} else {
				filePath = args[0]
				if charCount, err = strconv.ParseUint(args[2], 10, 64); err != nil {
					return err.Error()
				}
			}

			var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
			var f *os.File

			b := make([]rune, charCount)
			for i := range b {
				b[i] = letters[rand.Intn(len(letters))]
			}

			if f, err = os.Create(filePath); err != nil {
				return err.Error()
			}

			if _, err = f.WriteString(string(b)); err != nil {
				return err.Error()
			}

			return "Directory created."
		},
	},
	"openapp": &Event{
		Description:    "Opens an app, could be in the background.",
		ArgDescription: "<app name> <background flag>",
		Fn: func(args ...string) string {
			var inBackground bool

			if len(args) > 1 && args[1] != "0" {
				inBackground = true
			} else if len(args) > 0 {
				inBackground = false
			} else {
				return "Not enough arguments."
			}

			var cmd *exec.Cmd
			if inBackground {
				cmd = exec.Command("open", "-a", args[0], "-g")
			} else {
				cmd = exec.Command("open", "-a", args[0])
			}

			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				return err.Error()
			}

			if list := out.String(); err != nil {
				return list
			}

			return "App opened."
		},
	},
	"closeapp": &Event{
		Description:    "Closes an app.",
		ArgDescription: "<app name>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				closeApp(args[0])
			}

			return "App closed."
		},
	},
	"brightness": &Event{
		Description:    "Adjusts brightness level.",
		ArgDescription: "<brightness level 0 - 1; ex: 0.3>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				if output, err := storedActionScript("brightness.applescript", args[0]); err != nil {
					return err.Error()
				} else {
					return output
				}
			}
		},
	},
	"alert": &Event{
		Description:    "Generates OS X alert box.",
		ArgDescription: "<body> <title> <icon path> <button_1_text> <button_2_text>",
		Fn: func(args ...string) string {
			if len(args) < 6 {
				return "Not enough arguments."
			} else {
				if output, err := storedActionScript("alert.applescript", args[0], args[1], asPath(args[2]), args[3], args[4]); err != nil {
					return err.Error()
				} else {
					return output
				}
			}
		},
	},
	"volume": &Event{
		Description:    "Adjusts volume without UI",
		ArgDescription: "<volume % 0 - 100>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			} else {
				if output, err := actionScript(fmt.Sprintf("set volume output volume %s --100%", args[0])); err != nil {
					return err.Error()
				} else {
					return output
				}
			}
		},
	},
	"startaudio": &Event{
		Description:    "Starts playing audio file.",
		ArgDescription: "<absolute path to audio>",
		Fn: func(args ...string) string {
			if len(args) == 0 {
				return "Not enough arguments."
			}

			if *StartRepl {
				go termCommand("afplay", args[0])
			} else {
				if out, err := termCommand("afplay", args[0]); err != nil {
					return err.Error()
				} else if out != "" {
					return out
				}
			}

			return "Audio started"
		},
	},
	"stopaudio": &Event{
		Description:    "Stops audio.",
		ArgDescription: "",
		Fn: func(args ...string) string {
			var out string
			var err error
			if out, err = termCommand("killall", "afplay"); err != nil {
				return err.Error()
			}

			return out
		},
	},
}

func ShowCommands() string {
	var output string
	for cmd, info := range EventList {
		output = output + fmt.Sprintf("%s %s -- %s\n", cmd, info.ArgDescription, info.Description)
	}

	return output
}

type Event struct {
	Description     string                 // Name of event for logging.
	ArgDescription  string                 // A description of what arguments it takes.
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
	cmd := exec.Command(args[0], args[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	list := out.String()
	return list, nil
}

func closeApp(appName string) {
	actionScript(fmt.Sprintf("quit app \"%s\"", appName))
}

func asPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = "Macintosh HD" + path
	}

	return strings.Replace(path, "/", ":", -1)
}
