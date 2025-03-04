package main

import (
	"flag"
	"fmt"
	"slices"
)

func main() {
	validCmds := []string{"add", "list", "done"}
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

	fmt.Println(args[0])
}
