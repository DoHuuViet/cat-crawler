
package main

import "flag"
//import "fmt"
import "os"

import log "github.com/dmuth/google-go-log4go"


//
// Configuration for what was passed in on the command line.
//
type Config struct {
	SeedUrl string
}


/**
* Parse our command line arguments.
* @return {config} Our configuration info
*/
func ParseArgs() (retval Config) {

	retval = Config{""}

	flag.StringVar(&retval.SeedUrl, "seed_url", 
		"http://www.cnn.com/",
		"URL to start with.")
	h := flag.Bool("h", false, "To get this help")
	help := flag.Bool("help", false, "To get this help")
	debug_level := flag.String("debug-level", "info", "Set the debug level")

	flag.Parse()

	log.SetLevelString(*debug_level)
	log.Error("Debug level: " + *debug_level)

	if (*h || *help) {
		flag.PrintDefaults()
		os.Exit(1)
	}


	return(retval)

} // End of ParseArgs()

