package main

// CLI is not defined as OOPs as I thing func prog is good here.
// reduces the amount of code. and there is nothing mutable here to make an object.
import (
	"bufio"
	"log"
	"os"
	// string
	"strconv"
	"strings"
	// models
	logs "parking_lot/log"
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

	db := models.NewStoreyRunTimeDB(0)
	stdioLogger := logs.NewStdioLog()

	firstLine := true
	for scanner.Scan() {
		if firstLine {
			text := scanner.Text()
			command := parseCommand(text)
			if command[0] != models.CmdCreateParkingLot {
				panic("first command needs to be creating the storey")
			}
			maxSlots, err := strToInt(command[1])
			if err != nil {
				panic(err.Error())
			}
			// convert this to a new storey addition or update max slot method
			db = models.NewStoreyRunTimeDB(maxSlots)
			firstLine = false
			continue
		}
		stdioLogger.Log(processCommand(db, parseCommand(scanner.Text())))
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

// processCommand process each command
func processCommand(db models.DataStore, command []string) (models.StoreyResponse, error) {
	switch command[0] {
	case models.CmdPark:
		return db.Park(command[1], command[2])
	case models.CmdCreateParkingLot:
	case models.CmdStatus:
	case models.CmdLeave:
		slotPosition, err := strToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		return db.LeaveByPosition(slotPosition)
	case models.CmdRegistrationNumberByColor:
	case models.CmdSlotnoByCarColor:
	case models.CmdSlotnoByRegNumber:
	default:
	}

	return models.StoreyResponse{}, nil
}

// strToInt conver string to integer
func strToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
