package main

import (
	"math/rand"
	"regexp"
	"strings"
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

	rk := rand.Intn(len(keys) - 1)
	return this.GetCmd(keys[rk])
}

func (this *Profile) GetCmd(event string) string {
	var command string

	// Get a random command call
	if val, ok := this.EventData[event]; ok {
		command = event + " " + strings.TrimSpace(val[rand.Intn(len(val)-1)])
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
