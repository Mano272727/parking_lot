package main

import (
	"parking_lot/models"
)

func main() {
	models.NewStorey(4)
	models.NewCar("", "")
	ExecuteFile("samples/file_input.txt")
}
