package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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
var Command *string

func main() {
	StartRepl = flag.Bool("repl", false, "Start an interactive repl for command testing.")
	StartService = flag.String("service", "", "Start the service with the given profile JSON config file.")
	ResourceLocation = flag.String("resources", "/Applications/ComputerRoommate/Contents/Resources/", "The location of the resource files.")
	Command = flag.String("command", "", "A run and done command.")
	flag.Parse()

	if *StartRepl {
		go watchPs()
		repl()
	} else if *StartService != "" {
		var conf []byte
		var err error

		if _, err = os.Stat(*StartService); os.IsNotExist(err) {
			fmt.Println("Service profile configuration file not found.")
			os.Exit(1)
		}

		if conf, err = ioutil.ReadFile(*StartService); err != nil {
			fmt.Println("Service profile configuration file unreadable.")
			os.Exit(1)
		}

		var profile *Profile = &Profile{}
		if err := json.Unmarshal(conf, &profile); err != nil {
			fmt.Println("Unable to parse service config file.")
			os.Exit(1)
		}

		fmt.Println("Starting Computer Roommate service as " + profile.Name)

		go watchPs()
		service(profile, 6) // Hard coded to fire once a minute.
	} else if *Command != "" {
		parseCommand(*Command)
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

		if !parseCommand(line) {
			fmt.Println("Invalid command. Try \"commands\" for a command list.")
		}
	}
}

func parseCommand(command string) bool {
	r := regexp.MustCompile("'.+'|\".+\"|\\S+")
	args := r.FindAllString(command, -1)
	for k, v := range args {
		args[k] = strings.Replace(v, "\"", "", -1)
	}

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
