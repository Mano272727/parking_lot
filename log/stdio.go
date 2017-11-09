package log

import (
	"fmt"
)

type stdio struct{}

// NewStdioLog - creates a stdio logger
func NewStdioLog() *stdio {
	return &stdio{}
}

// Log - logs the output, this logger logs to stdio.
func (s *stdio) Log(loggable Loggable, err error) error {
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Println(loggable)
	return nil
}
