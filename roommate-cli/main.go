package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

var StartRepl *bool
var ListProfiles *bool
var StartService *string
var ResourceLocation *string

func main() {
	StartRepl = flag.Bool("repl", false, "Start an interactive repl for command testing.")
	ListProfiles = flag.Bool("list-profiles", false, "Lists all roommate profiles.")
	StartService = flag.String("service", "", "Start the service with the given profile name.")
	ResourceLocation = flag.String("resources", "/Applications/ComputerRoommate/Contents/Resources/", "The location of the resource files.")
	flag.Parse()

	if *StartRepl {
		go watchPs()
		repl()
	} else if *ListProfiles {
		profiles := make([]string, len(ProfileList))
		for k, v := range ProfileList {
			profiles = append(profiles, k+" -- "+v.Description)
		}
		fmt.Println("Roommate Profiles:" + strings.Join(profiles, "\n"))
		os.Exit(0)
	} else if *StartService != "" {
		profiles := make(map[string]bool, len(ProfileList))
		for k, _ := range ProfileList {
			profiles[k] = true
		}

		if _, ok := profiles[*StartService]; !ok {
			fmt.Println("Invalid profile selected. Try -list-profiles")
			os.Exit(1)
		}

		go watchPs()
		service(ProfileList[*StartService], 6) // Hard coded to fire once a minute.
	} else if len(os.Args) > 1 {
		args := os.Args[1:]
		parseCommand(args)
		os.Exit(0)
	} else {
		fmt.Println("You didn't tell me to do anything. Try -help\n")
		os.Exit(1)
	}
}

func service(roommate *Profile, top int) {
	fmt.Println("Starting super annoying service in the background.")

	go func() {
		for {
			if rand.Intn(top) == 1 {
				go parseCommand(roommate.GetRandCmd())
			}

			time.Sleep(1000 * time.Millisecond * 10) // 10s sleep
		}
	}()
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
