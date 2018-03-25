package main

// CLI is not defined as OOPs as I thing func prog is good here.
// reduces the amount of code. and there is nothing mutable here to make an object.
import (
	"bufio"
	"fmt"
	"log"
	"os"
	// string
	"strconv"
	"strings"
	// models
	logs "github.com/sumitasok/parking_lot/log"
	"github.com/sumitasok/parking_lot/models"
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
			stdioLogger.Log(models.NewDbResponse(*db.Storeys[0], models.CmdCreateParkingLot), nil)
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

// InteractiveSession take the user through interactive session.
func InteractiveSession() error {
	command := "Start"
	stdioLogger := logs.NewStdioLog()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nInput")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	commands := parseCommand(text)
	if commands[0] != models.CmdCreateParkingLot {
		panic("first command needs to be creating the storey")
	}
	maxSlots, err := strToInt(commands[1])
	if err != nil {
		panic(err.Error())
	}
	// convert this to a new storey addition or update max slot method
	db := models.NewStoreyRunTimeDB(maxSlots)
	fmt.Println("\nOutput")
	stdioLogger.Log(models.NewDbResponse(*db.Storeys[0], models.CmdCreateParkingLot), nil)

	for command != "Exit" {
		fmt.Println("\nInput")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		// text = text[:len(text)-1]
		commands := parseCommand(text)
		response, err := processCommand(db, commands)
		fmt.Println("\nOutput")
		stdioLogger.Log(response, err)
		command = commands[0]
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
	case models.CmdCreateParkingLot:
		maxSlots, err := strToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		return db.AddStorey(maxSlots)
	case models.CmdPark:
		return db.Park(command[1], command[2])
	case models.CmdCreateParkingLot:
	case models.CmdStatus:
		return db.All()
	case models.CmdLeave:
		slotPosition, err := strToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		return db.LeaveByPosition(slotPosition)
	case models.CmdRegistrationNumberByColor:
		return db.FindAllByColor(command[1], models.CmdRegistrationNumberByColor)
	case models.CmdSlotnoByCarColor:
		return db.FindAllByColor(command[1], models.CmdSlotnoByCarColor)
	case models.CmdSlotnoByRegNumber:
		return db.FindByRegistrationNumber(command[1])
	default:
	}

	return models.StoreyResponse{}, nil
}

// strToInt conver string to integer
func strToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
