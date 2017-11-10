package main

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func ExampleExecutableFile() {
	ExecuteFile("samples/file_input.txt")
	// without fmt the exmaple is not working. Investigation pending.
	fmt.Println("")
	// output:
	// Created   a   parking   lot   with   6   slots
	// Allocated   slot   number:   1
	// Allocated   slot   number:   2
	// Allocated   slot   number:   3
	// Allocated   slot   number:   4
	// Allocated   slot   number:   5
	// Allocated   slot   number:   6
	// Slot number 4 is free
	// Slot   No. Registration No	Color
	// 1 KA-01-HH-1234 White
	// 2 KA-01-HH-9999 White
	// 3 KA-01-BB-0001 Black
	// 5 KA-01-HH-2701 Blue
	// 6 KA-01-HH-3141 Black
	// Allocated   slot   number:   4
	// Sorry,   parking   lot   is   full
	// KA-01-HH-1234,   KA-01-HH-9999,   KA-01-P-333
	// 1,   2,   4
	// 6
	// Not   found
	//
}

func TestParseCommand(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"create_parking_lot", "6"}, parseCommand("create_parking_lot   6"))
	assert.Equal([]string{"park", "KA-01-HH-1234", "White"}, parseCommand("park   KA-01-HH-1234	White"))

	assert.True(true)
}
