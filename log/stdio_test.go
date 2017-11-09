package log

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

type LoggableSample struct {
}

func (l LoggableSample) String() string {
	return fmt.Sprintf("loggable samples")
}

// This test only asserts that the operation completes.
// To-Do: The writing to stdio testing
// https://stackoverflow.com/questions/26804642/how-to-test-a-functions-output-stdout-stderr-in-go-unit-tests
// https://stackoverflow.com/questions/10473800/in-go-how-do-i-capture-stdout-of-a-function-into-a-string/10476304#10476304
func TestLogStdio(t *testing.T) {
	assert := assert.New(t)

	logObj := NewStdioLog()
	logObj.Log(LoggableSample{}, nil)

	assert.True(true)
}
