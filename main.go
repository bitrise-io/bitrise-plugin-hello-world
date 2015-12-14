package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

//-----------------------------------
// Communication
//-----------------------------------

// Plugin output (logs, errors, ...) should send to os.Stderr
// Messages to bitrise (plugin requirements, plugin version, ...) should send to os.Stdout as JSON string

// Print errors to os.Stderr
func sendFatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", msg)
	os.Exit(1)
}

// Print outputs to os.Stderr
func sendMessageToOutput(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", msg)
}

// Send message to bitrise througth os.Stdout as JSON string
func sendMessageToBitrise(format string, a ...interface{}) {
	msgStr := fmt.Sprintf(format, a...)
	msg := map[string]string{
		"bitrise-info": msgStr,
	}

	msgJSON, err := json.Marshal(msg)
	if err != nil {
		sendFatal("%s", err)
	}

	fmt.Fprintf(os.Stdout, "%s\n", msgJSON)
}

//-----------------------------------
// CLI
//-----------------------------------

// Command model
type Command struct {
	Name string
	Func func() error
}

// Available commands
var commands = []Command{
	Command{
		Name: "requirements",
		Func: requirements,
	},
}

// Requirements command, all plugin should implement this.
// After installing a plugin, bitrise ask a plugin for plugin requirements.
func requirements() error {
	type Requirement struct {
		ToolName   string
		MinVersion string
		MaxVersion string
	}

	var requirements = []Requirement{
		Requirement{
			ToolName:   "bitrise",
			MinVersion: "1.3.0",
			MaxVersion: "",
		},
		Requirement{
			ToolName:   "stepman",
			MinVersion: "1.3.0",
			MaxVersion: "",
		},
		Requirement{
			ToolName:   "envman",
			MinVersion: "1.3.0",
			MaxVersion: "",
		},
	}

	bytes, err := json.Marshal(requirements)
	if err != nil {
		return err
	}

	sendMessageToBitrise(string(bytes))

	return nil
}

// Bitrise send messages to Plugin on os.Stdin as JSON string.
// Message can be request for command (ex: requirements, doSomething, ...),
// or informations from bitrise (ex: bitrise version, bitrise workflow path, ...)
func parseCli() (string, error) {
	args := os.Args
	if len(args) < 2 {
		return "", errors.New("Invalid plugin usage, no bitrise infos, even command passed to the plugin")
	}

	msg := os.Args[1]
	return msg, nil
}

//-----------------------------------
// Main
//-----------------------------------

func main() {
	message, err := parseCli()
	if err != nil {
		sendFatal("Failed to parse cli, err: %s", err)
	}

	messageIsCommand := false
	for _, command := range commands {
		if command.Name == message {
			if err := command.Func(); err != nil {
				sendFatal("Failed to perform command (%s), err: %s", command.Name, err)
			}
			messageIsCommand = true
		}
	}

	if !messageIsCommand {
		sendMessageToOutput("Bitrise message: %s", message)
		sendMessageToOutput("Hello World!")
	}
}
