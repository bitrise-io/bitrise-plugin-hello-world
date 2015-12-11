package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func printFatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", msg)
	os.Exit(1)
}

func printHello(name string) {
	msg := fmt.Sprintf("Hello %s!", name)
	fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", msg)
}

func printBitriseInfo() {
	msg := map[string]string{
		"bitrise-info": "Bitrise CLI plugin sample",
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		printFatal("%s", err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", msgJSON)
}

func parseCli(args []string) (string, error) {
	if len(args) > 2 {
		return "", fmt.Errorf("Only one argument allowed, provided (%d)", len(os.Args)-1)
	}

	name := "Bitriser"
	if len(args) == 2 {
		name = args[1]
	}
	return name, nil
}

func main() {
	// Parse cli
	name, err := parseCli(os.Args)
	if err != nil {
		printFatal("Failed to parse cli, err: %s", err)
	}

	printHello(name)
	printBitriseInfo()
}
