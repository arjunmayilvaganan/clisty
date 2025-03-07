package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

const (
	Add                  = "add"
	List                 = "list"
	Done                 = "done"
	Help                 = "help"
	DefaultStoreFilePath = "store.json"
	Capacity             = 10000
)

var config Configuration
var store Store

type Configuration struct {
	StoreFilePath string
}

type Item struct {
	Id   uint64   `json:"id"`
	Name string   `json:"name"`
	Done bool     `json:"done"`
	Tags []string `json:"tags"`
}

type Store struct {
	Capacity        uint64 `json:"capacity"`
	TotalLength     uint64 `json:"totalLength"`
	CompletedLength uint64 `json:"completedLength"`
	Items           []Item `json:"items"`
}

func (s *Store) addItem(item Item) {
	if s.Capacity < s.TotalLength+1 {
		fmt.Println("The store is already at maximum capacity")
		return
	}

	if len(s.Items) == 0 {
		s.Items = []Item{}
	}
	item.Id = s.TotalLength + 1
	s.Items = append(s.Items, item)
	s.TotalLength = s.TotalLength + 1
	s.listCompletedItems()
}

func (s *Store) listCompletedItems() {
	s.listItems(false)
}

func (s *Store) listItems(listCompleted bool) {
	for _, item := range s.Items {
		if !item.Done {
			fmt.Println("[ ]", item.Id, item.Name)
		}

		if item.Done && listCompleted {
			fmt.Println("[X]", item.Id, item.Name)
		}
	}
}

func (s *Store) markDone(item Item) {
	idx := slices.IndexFunc(s.Items, func(i Item) bool { return i.Id == item.Id })
	s.Items[idx].Done = true
	s.CompletedLength = s.CompletedLength + 1
	s.listCompletedItems()
}

func parseAddArgs(args []string) Item {
	return Item{
		Name: args[0],
		Done: false,
	}
}

func parseDoneArgs(args []string) Item {
	id, err := strconv.ParseUint(args[0], 10, 64)
	if err != nil {
		log.Fatalln(err)
	}
	return Item{Id: id}
}

func initialize(storeFilePath string) {
	storeData, err := os.ReadFile(storeFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	config = Configuration{StoreFilePath: storeFilePath}
	json.Unmarshal(storeData, &store)
	store.Capacity = Capacity
}

func shutdown(storeFilePath string) {
	storeData, err := json.Marshal(store)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(storeFilePath, storeData, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func displayHelp() {
	fmt.Println("Usage: clisty [OPTIONS]")
	fmt.Println("Options:")
	fmt.Println("  clisty add \"<To-do item description goes here>\"\t-- To add a new item to the list")
	fmt.Println("  clisty done \"<ID of the item to be marked done>\"\t-- To mark that item as done")
	fmt.Println("  clisty list\t\t\t\t\t\t-- Prints all pending items from the list")
	fmt.Println("  clisty list all\t\t\t\t\t-- Prints both pending and completed items from the list")
	fmt.Println("  clisty help\t\t\t\t\t\t-- Prints this usage guide again")
}

func main() {
	initialize(DefaultStoreFilePath)
	validCmds := []string{Add, List, Done, Help}
	args := os.Args[1:]
	if len(args) < 1 {
		displayHelp()
		return
	}

	if !slices.Contains(validCmds, args[0]) {
		displayHelp()
		return
	}

	switch args[0] {
	case Add:
		item := parseAddArgs(args[1:])
		store.addItem(item)
	case List:
		listAllItems := len(args) > 1 && args[1] == "all"
		store.listItems(listAllItems)
	case Done:
		item := parseDoneArgs(args[1:])
		store.markDone(item)
	case Help:
		displayHelp()
	}
	shutdown(DefaultStoreFilePath)
}
