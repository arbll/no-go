package main

import (
	"github.com/Omen-/no-go/web"
	"github.com/Omen-/no-go/config"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
	}

	config.Load(args[0])
	web.Run()
}

func printUsage()  {

}