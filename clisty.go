package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := flag.String("cmd", "", "The default command supplied to the executable")
	flag.Parse()
	if *cmd != "" {
		fmt.Println(*cmd)
	} else {
		fmt.Println("Nothing was entered!")
	}
}
