package main

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type Profile struct {
	Name        string              `json:"name"`
	Description string              `json:"description"`
	Options     map[string]string   `json:"options"`
	EventData   map[string][]string `json:"events"`
}

func (this *Profile) GetRandCmd() string {
	keys := make([]string, 0, len(this.EventData))
	for k := range this.EventData {
		keys = append(keys, k)
	}

	rand.Seed(time.Now().Unix())
	rk := rand.Intn(len(keys))
	return this.GetCmd(keys[rk])
}

func (this *Profile) GetCmd(event string) string {
	var command string

	// Get a random command call
	if val, ok := this.EventData[event]; ok {
		if len(val) == 1 {
			command = event + " " + strings.TrimSpace(val[0])
		} else {
			rand.Seed(time.Now().Unix())
			command = event + " " + strings.TrimSpace(val[rand.Intn(len(val))])
		}
	} else {
		return ""
	}

	// Scan it for variables
	for search, replace := range this.Options {
		re := regexp.MustCompile("\\[" + search + "\\]")
		command = re.ReplaceAllString(command, replace)
	}

	// Return it
	return command
}
