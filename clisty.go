package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
)

const (
	Add                  = "add"
	List                 = "list"
	Done                 = "done"
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
	fmt.Println(s.Items)
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
		panic(err)
	}
	return Item{Id: id}
}

func initialize(storeFilePath string) {
	storeData, err := os.ReadFile(storeFilePath)
	if err != nil {
		panic(err)
	}
	config = Configuration{StoreFilePath: storeFilePath}
	json.Unmarshal(storeData, &store)
	store.Capacity = Capacity
}

func shutdown(storeFilePath string) {
	storeData, err := json.Marshal(store)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(storeFilePath, storeData, 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	initialize(DefaultStoreFilePath)
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

	switch args[0] {
	case Add:
		item := parseAddArgs(args[1:])
		store.addItem(item)
	case List:
		store.listCompletedItems()
	case Done:
		item := parseDoneArgs(args[1:])
		store.markDone(item)
	}
	shutdown(DefaultStoreFilePath)
}
