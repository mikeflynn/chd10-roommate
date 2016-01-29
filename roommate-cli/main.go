package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	//"math"
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var StartRepl *bool
var StartService *bool

func main() {
	StartRepl = flag.Bool("repl", false, "Start an interactive repl for command testing.")
	StartService = flag.Bool("service", false, "Start the app watching and random task service.")
	flag.Parse()

	if *StartRepl {
		go watchPs()
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

	// Random Event Loop
	go func() {
		for {
			//ts := time.Now().UnixNano()

			time.Sleep(1000 * time.Millisecond)
		}
	}()

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
		if !parseCommand(args) {
			fmt.Println("Invalid command. Try \"commands\" for a command list.")
		}
	}
}

func parseCommand(args []string) bool {
	if args[0] == "commands" {
		fmt.Println(ShowCommands())
	} else if args[0] == "quit" || args[0] == "exit" {
		os.Exit(0)
	} else if event, ok := EventList[args[0]]; ok {
		event.Run(args[1:]...)
	} else {
		return false
	}

	return true
}

func watchPs() {
	found := map[string]bool{}

	for {
		cmd := exec.Command("ps", "-x")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		list := out.String()

		for _, app := range WatchList {
			if app.Timeout < time.Now().Unix() {
				delete(found, app.Name)
				continue
			}

			if matched, err := regexp.MatchString("MacOS/"+app.Name+"\\b", list); err == nil && matched {
				if _, ok := found[app.Name]; !ok {
					found[app.Name] = true
					go func() {
						time.Sleep(1000 * time.Millisecond * 1)

						closeApp(app.Name)

						// Payload
						if len(app.Payload) > 0 {
							parseCommand(app.Payload)
						}

						delete(found, app.Name)
					}()
				}
			}
		}
	}
}
