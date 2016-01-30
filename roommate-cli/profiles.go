package main

import (
	"math/rand"
	"regexp"
	"strings"
)

var ProfileList map[string]*Profile = map[string]*Profile{
	"autumn": {
		Description: "A stoner.",
		Options: map[string]string{
			"icon": "...",
			"name": "Autumn",
		},
		EventData: map[string][]string{
			"wallpaper": []string{
				"...img 1",
				"...img 2",
			},
		},
	},
}

type Profile struct {
	Description string
	Options     map[string]string
	EventData   map[string][]string
}

func (this *Profile) GetRandCmd() []string {
	keys := make([]string, 0, len(this.EventData))
	for k := range this.EventData {
		keys = append(keys, k)
	}

	rk := rand.Intn(len(keys) - 1)
	return this.GetCmd(keys[rk])
}

func (this *Profile) GetCmd(event string) []string {
	command := []string{event}

	// Get a random command call
	if val, ok := this.EventData[event]; ok {
		command = append(command, strings.Split(strings.TrimSpace(val[rand.Intn(len(val)-1)]), " ")...)
	} else {
		return []string{}
	}

	// Scan it for variables
	for idx, part := range command {
		for search, replace := range this.Options {
			re := regexp.MustCompile("\\[" + search + "\\]")
			command[idx] = re.ReplaceAllString(part, replace)
		}
	}

	// Return it
	return command
}
