package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	//"math"
	"os"
	"strings"
)

var StartRepl *bool
var StartService *bool

func main() {
	StartRepl = flag.Bool("repl", false, "Start an interactive repl for command testing.")
	StartService = flag.Bool("service", false, "Start the app watching and random task service.")
	flag.Parse()

	if *StartRepl {
		repl()
	} else if *StartService {
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
	if args[0] == "commands" {
		fmt.Println(ShowCommands())
	} else if args[0] == "quit" {
		os.Exit(0)
	} else if event, ok := EventList[args[0]]; ok {
		event.Run(args[1:]...)
	} else {
		fmt.Println("Invalid command. Try \"commands\" for a command list.")
	}
}

func watchPs() {
	//found := map[string]bool{}

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
	/*
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
	*/
}
