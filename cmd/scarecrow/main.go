/*
Main entry point for the Scarecrow chatbot application.
*/
package main

import (
	"flag"
	"fmt"
	scarecrow "github.com/aichaos/scarecrow/src"
	_ "github.com/aichaos/scarecrow/src/listeners/console"
	_ "github.com/aichaos/scarecrow/src/listeners/slack"
	_ "github.com/aichaos/scarecrow/src/listeners/xmpp"
	"os"
)

func main() {
	// Collect command line parameters.
	debug := flag.Bool("debug", false, "Enable debug logging.")
	version := flag.Bool("version", false, "Show the version number and exit.")
	flag.Parse()

	if *version == true {
		fmt.Printf("This is Scarecrow, version %s\n", scarecrow.VERSION)
		os.Exit(0)
	}

	// Create the bot instance.
	bot := scarecrow.New()
	bot.SetDebug(*debug)

	bot.Start()
}
