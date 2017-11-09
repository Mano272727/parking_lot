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

func TestStorey_Park(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	storey.slotList.UpdatePosition(2)
	storey.Park("numberPlate - x", "color - x")

	assert.Equal("numberPlate - x", storey.slotList.car.numberPlate)
	assert.Equal(1, storey.slotList.Position())

	assert.True(true)
}

func TestSlot_AddNext(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)

	assert.Equal(1, storey.slotList.Position())
	assert.Equal(2, storey.OccupancyCount())
	assert.Equal(2, storey.slotList.nextSlot.Position())

	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())

	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)
	assert.Equal(2, storey.slotList.nextSlot.Position())
	assert.Equal(3, storey.slotList.nextSlot.nextSlot.Position())

	assert.Equal("numberPlate - y", storey.slotList.nextSlot.car.numberPlate)
	assert.Equal("numberPlate - x", storey.slotList.nextSlot.nextSlot.car.numberPlate)

	assert.True(true)
}

func TestNewCar(t *testing.T) {
	assert := assert.New(t)

	car := NewCar("KL-00-0000", "red")
	assert.NotEmpty(car)

	assert.True(true)
}
