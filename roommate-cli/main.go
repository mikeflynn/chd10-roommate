package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	//"math"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/deckarep/gosx-notifier"
)

var startRepl *bool
var startService *bool

func main() {
	startRepl = flag.Bool("repl", false, "Start an interactive repl for command testing.")
	startService = flag.Bool("service", false, "Start the app watching and random task service.")
	flag.Parse()

	if *startRepl {
		repl()
	} else if *startService {
		service()
	} else if len(os.Args) > 1 {
		args := os.Args[1:]
		parseCommand(args)
	} else {
		fmt.Println("You didn't tell me to do anything. Try -help\n")
	}
}

func service() {
	fmt.Println("Starting super annoying service in the background.")
	watchPs()
}

func repl() {
	var line string
	var err error

	fmt.Printf("Computer Roommate Terminal\n" +
		"==========================\n" +
		"Type \"commands\" for a list of commands.\n\n")

	for {
		fmt.Printf("> ")

		in := bufio.NewReader(os.Stdin)
		if line, err = in.ReadString('\n'); err != nil {
			log.Fatal(err)
		}

		args := strings.Split(strings.TrimSpace(line), " ")
		parseCommand(args)
	}
}

func parseCommand(args []string) {
	var err error

	switch {
	case args[0] == "notify":
		n := &Notification{
			Body:     "Can you pick some up?",
			Title:    "Hey...",
			Subtitle: "We're out of your milk.",
			Image:    "/Applications/ComputerRoommate.app/Contents/Resources/icon.ico",
		}

		err := n.notify()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Notification sent.")
	case args[0] == "wallpaper" && len(args) > 1:
		changeWallpaper(args[1])

		fmt.Println("Wallpaper updated.")
	case args[0] == "openapp" && len(args) > 1:
		if len(args) > 2 && args[2] != "0" {
			openApp(args[1], true)
		} else {
			openApp(args[1], false)
		}

		fmt.Println("App opened.")
	case args[0] == "closeapp" && len(args) > 1:
		closeApp(args[1])

		fmt.Println("App closed.")
	case args[0] == "quicklook" && len(args) > 1:
		_, err := quickLook(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Quicklook popped.")
	case args[0] == "startaudio" && len(args) > 1:
		if *startRepl {
			go func() {
				if out, err := startAudio(args[1]); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(out)
				}
			}()
		} else {
			if out, err := startAudio(args[1]); err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(out)
			}
		}

		fmt.Println("Audio playing.")
	case args[0] == "stopaudio":
		_, err := stopAudio()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Audio stopped.")
	case args[0] == "openfile" && len(args) > 1:
		_, err := openFile(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("That thing opened.")
	case args[0] == "makedir" && len(args) > 1:
		if err := createDir(args[1]); err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Directory created.")
	case args[0] == "makefile" && len(args) > 1:
		var count uint64 = 0
		if len(args) > 2 {
			if count, err = strconv.ParseUint(args[2], 10, 64); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if err := createFile(args[1], count); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Directory created.")
		}
	case args[0] == "movefile" && len(args) > 2:
		if err := moveFile(args[1], args[2]); err != nil {
			fmt.Println(err.Error())
		}
	case args[0] == "brightness" && len(args) > 1:
		if output, err := storedActionScript("brightness.applescript", args[1]); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	case args[0] == "alert" && len(args) > 5:
		if output, err := storedActionScript("alert.applescript", args[1], args[2], asPath(args[3]), args[4], args[5]); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	case args[0] == "volume" && len(args) > 1:
		if _, err := actionScript(fmt.Sprintf("set volume output volume %s --100%", args[1])); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Volume set to " + args[1])
		}
	case args[0] == "commands":
		fmt.Println("Valid commands:\n" +
			"notify\n" +
			"wallpaper <absolute path to image>\n" +
			"quicklook <absolute path to image>\n" +
			"movefile <absolute path to source> <absolute path to desitination>\n" +
			"openfile <absolute path to file>\n" +
			"makedir <absolute path to new directory>\n" +
			"makefile <absolute path to new file> <number of chars to fill it with>\n" +
			"openapp <app name> <background flag>\n" +
			"closeapp <app name>\n" +
			"brightness <brightness level 0 - 1; ex: 0.3>\n" +
			"alert <body> <title> <icon path> <button_1_text> <button_2_text>\n" +
			"volume <volume % 0 - 100>\n" +
			"startaudio <absolute path to audio>\n" +
			"stopaudio\n")
	case args[0] == "quit":
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Invalid command. Try \"commands\" for a command list.")
	}
}

func asPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = "Macintosh HD" + path
	}

	return strings.Replace(path, "/", ":", -1)
}

func watchPs() {
	found := map[string]bool{}

	/*
		go func() {
			for {
				ts := time.Now().UnixNano()
				if math.Mod(float64(ts), 15) == 0.0 {
					changeWallpaper("/Applications/ComputerRoommate.app/Contents/Resources/wallpaper.png")
				} else if math.Mod(float64(ts), 9) == 0.0 {
					fmt.Println("Making files...")
					//createFile("DO NOT TOUCH MY STUFF"+strconv.FormatInt(ts, 10)+".txt", 1000)
				} else if math.Mod(float64(ts), 3) == 0.0 {
					//openApp("Messages", true)
					//time.Sleep(2000 * time.Millisecond)
					//closeApp("Messages")
				}

				time.Sleep(1000 * time.Millisecond)
			}
		}()
	*/

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
			name := "iTunes"
			if _, ok := found[name]; !ok {
				found[name] = true
				go func() {
					time.Sleep(1000 * time.Millisecond * 1)

					closeApp("iTunes")

					if output, err := storedActionScript("alert.applescript", "I'm in here!", "Ugggghhh!", asPath("/Applications/ComputerRoommate.app/Contents/Resources/icon.ico"), "Ok", "Hurry up!"); err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println(output)
					}

					// Close it.
					delete(found, name)
				}()
			}
		}

		if matched, err := regexp.MatchString("MacOS/Safari\\b", list); err == nil && matched {
			name := "Safari"
			if _, ok := found[name]; !ok {
				found[name] = true
				go func() {
					time.Sleep(1000 * time.Millisecond * 1)

					closeApp("Safari")

					if output, err := storedActionScript("alert.applescript", "I'll be out in a minute!", "Please Knock!", asPath("/Applications/ComputerRoommate.app/Contents/Resources/icon.ico"), "Ok", "Hurry up!"); err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println(output)
					}

					// Close it.
					delete(found, name)
				}()
			}
		}

		if matched, err := regexp.MatchString("MacOS/Keynote\\b", list); err == nil && matched {
			name := "Keynote"
			if _, ok := found[name]; !ok {
				found[name] = true
				go func() {
					time.Sleep(1000 * time.Millisecond * 1)

					closeApp("Keynote")

					if output, err := storedActionScript("alert.applescript", "Give me a second!", "Ugggghhh!", asPath("/Applications/ComputerRoommate.app/Contents/Resources/icon.ico"), "Ok", "Hurry up!"); err != nil {
						fmt.Println(err.Error())
					} else {
						fmt.Println(output)
					}

					// Close it.
					delete(found, name)
				}()
			}
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

func createFile(filePath string, charCount uint64) error {
	var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var f *os.File
	var err error

	b := make([]rune, charCount)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	if f, err = os.Create(filePath); err != nil {
		return err
	}

	if _, err = f.WriteString(string(b)); err != nil {
		return err
	}

	return nil
}

func createDir(dirPath string) error {
	return os.MkdirAll(dirPath, 0777)
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

func openApp(appName string, inBackground bool) (string, error) {
	var cmd *exec.Cmd
	if inBackground {
		cmd = exec.Command("open", "-a", appName, "-g")
	} else {
		cmd = exec.Command("open", "-a", appName)
	}

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

func closeApp(appName string) {
	actionScript(fmt.Sprintf("quit app \"%s\"", appName))
}

func moveFile(sourcePath string, destPath string) error {
	return os.Rename(sourcePath, destPath)
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
