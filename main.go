package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Print errors to os.Stderr
func printFatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", msg)
	os.Exit(1)
}

// Print outputs to os.Stderr
func printOutput(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", msg)
}

// Send infos to bitrise througth os.Stdout as JSON string
func sendBitriseInfo() {
	msg := map[string]string{
		"bitrise-info": "Bitrise CLI plugin sample",
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		printFatal("%s", err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", msgJSON)
}

// Bitrise send informations to Plugin on os.Stdin as JSON string
func parseBitriseInfo() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", errors.New("Invalid plugin usage, no bitrise infos passed to the plugin")
	}

	infos := args[1]
	return infos, nil
}

func main() {
	// Parse cli
	infos, err := parseBitriseInfo()
	if err != nil {
		printFatal("Failed to parse cli, err: %s", err)
	}

	printOutput("Hello World!")
	printOutput("Bitrise infos: %s", infos)
	sendBitriseInfo()
}
