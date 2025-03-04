package main

import (
	"flag"
	"fmt"
	"slices"
)

type Item struct {
	Id   uint64   `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

const (
	Add  = "add"
	List = "list"
	Done = "done"
)

func parseAddArgs(args []string) Item {
	fmt.Println(args)
	return Item{}
}

func addItem(item Item, items []Item) {
	fmt.Println(item)
	fmt.Println(items)
	fmt.Println("added")
}

func listItems(items []Item) {
	fmt.Println(items)
}

func parseDoneArgs(args []string) Item {
	fmt.Println(args)
	return Item{}
}

func markDone(item Item, items []Item) {
	fmt.Println(item)
	fmt.Println(items)
	fmt.Println("marked done")
}

func main() {
	validCmds := []string{Add, List, Done}
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("You need to provide a command")
		fmt.Println("The valid commands accepted are: ", validCmds)
		return
	}

	if !slices.Contains(validCmds, args[0]) {
		fmt.Println("Provided command is not valid")
		fmt.Println("The valid commands accepted are: ", validCmds)
		return
	}

	items := []Item{}

	switch args[0] {
	case Add:
		item := parseAddArgs(args[1:])
		addItem(item, items)
	case List:
		listItems(items)
	case Done:
		item := parseDoneArgs(args[1:])
		markDone(item, items)
	}
}
