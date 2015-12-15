package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//-----------------------------------
// Communication
//-----------------------------------

// Plugin output (logs, errors, ...) should send to os.Stderr
// Messages to bitrise (plugin requirements, plugin version, ...) should send to os.Stdout as JSON string

// Bitrise send message to Plugin througth BITRISE_PLUGINS_MESSAGE environment
func readMessageFromBitrise() (map[string]interface{}, error) {
	messageFromBitrise := os.Getenv("BITRISE_PLUGINS_MESSAGE")

	if messageFromBitrise != "" {
		var message map[string]interface{}
		if err := json.Unmarshal([]byte(messageFromBitrise), &message); err != nil {
			return map[string]interface{}{}, err
		}

		return message, nil
	}

	return map[string]interface{}{}, nil
}

// Print errors to os.Stderr
func sendFatal(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[31;1m%s\x1b[0m\n", msg)
	os.Exit(1)
}

// Print outputs to os.Stderr
func sendGreenMessageToOutput(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "\x1b[32;1m%s\x1b[0m\n", msg)
}

func sendMessageToOutput(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}

// Send message to bitrise througth os.Stdout as JSON string
func sendMessageToBitrise(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(os.Stdout, "%s\n", msg)
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
	Command{
		Name: "version",
		Func: version,
	},
}

// Requirements command, all plugin should implement this.
// After installing a plugin, bitrise ask a plugin for plugin requirements,
// this case plugin have to put answer on os.Stdout as JSON string
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
			MinVersion: "0.9.17",
			MaxVersion: "",
		},
		Requirement{
			ToolName:   "envman",
			MinVersion: "1.0.0",
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

// version command
// Normal plugin command, this case plugin have to put answer on os.Stderr
func version() error {
	scriptDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}
	versionPth := path.Join(scriptDir, "version")
	versionBytes, err := ioutil.ReadFile(versionPth)

	sendMessageToOutput(string(versionBytes))

	return nil
}

func parseCommandFromCli() string {
	args := os.Args
	if len(args) < 2 {
		return ""
	}

	return os.Args[1]
}

//-----------------------------------
// Main
//-----------------------------------

func main() {
	messageFromBitrise, err := readMessageFromBitrise()
	if err != nil {
		sendFatal("Failed to read message from bitrise, err: %s", err)
	}

	commandFromCli := parseCommandFromCli()

	commandFound := false
	for _, command := range commands {
		if command.Name == commandFromCli {
			if err := command.Func(); err != nil {
				sendFatal("Failed to perform command (%s), err: %s", command.Name, err)
			}
			commandFound = true
		}
	}

	if !commandFound {
		sendGreenMessageToOutput("Hello World!")
		sendMessageToOutput("Message from bitrise: %s", messageFromBitrise)
	}
}
