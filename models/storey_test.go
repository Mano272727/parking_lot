package models

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorey(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.NotEmpty(storey)

	assert.True(true)
}

func TestStorey_OccupancyCount(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	assert.True(true)
}

func TestNewCar(t *testing.T) {
	assert := assert.New(t)

	car := NewCar("KL-00-0000", "red")
	assert.NotEmpty(car)

	assert.True(true)
}
