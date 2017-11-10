package main

// CLI is not defined as OOPs as I thing func prog is good here.
// reduces the amount of code. and there is nothing mutable here to make an object.
import (
	"bufio"
	"log"
	"os"
	// string
	"strings"
	// models
	"parking_lot/models"
)

var (
	// CommandSeparator is the separator used by default in the input
	CommandSeparator = " "
	// Tab is the tab character.
	Tab = "\t"
)

// ExecuteFile taskes in a file path and execute the commands in the file.
func ExecuteFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processCommand(parseCommand(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// parseCommand takes in a command string and converts it to a string array
// by removing the tab character in the command.
func parseCommand(command string) []string {
	parsedCommand := []string{}

	// remove the tabs in between the command.
	command = strings.Replace(command, Tab, CommandSeparator, -1)

	// remove the empty string
	for _, s := range strings.Split(command, CommandSeparator) {
		if s != "" {
			parsedCommand = append(parsedCommand, s)
		}
	}

	return parsedCommand
}

func processCommand(command []string) (models.SlotResponse, error) {
	switch command[0] {
	case "create_parking_lot":
	default:
	}

	return models.SlotResponse, nil
}
