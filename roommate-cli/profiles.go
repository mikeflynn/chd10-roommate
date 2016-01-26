package main

type Roommate struct {
	Name       string
	Picture    string
	Attributes AttributeList
	EventData  map[string]map[string]string
}

type AttributeList struct {
	Clean  int
	Angry  int
	Creepy int
}
