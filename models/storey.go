package models

import (
	"errors"
)

var (
	// ErrMaxSlotReached - error when maximum number of slots allowed in a storey is reached.
	ErrMaxSlotReached = errors.New("Max slot reached")
)

// Storey - holds the slots and the details about the storey. (Multi storey
// can have many Storey objects)
type Storey struct {
	maxSlots int
	// Linked list should keep it memory efficient.
	slotList *Slot
}

// Park - check if the Slot is available
// if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Park(numberPlate, color string) (*Slot, error) {
	slot := &Slot{}
	if s.OccupancyCount() >= s.maxSlots {
		return slot, ErrMaxSlotReached
	}

	car := NewCar(numberPlate, color)

	if s.OccupancyCount() == 0 {
		slot := NewSlot(car, 0)
		s.slotList = slot
		return slot, nil
	}

	return slot, nil
}

// Leave - check if the Slot is available
// if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Leave(numberPlate string) (*Slot, error) {
	return &Slot{}, nil
}

// OccupancyCount returns the number of slots occupied in this storey.
func (s *Storey) OccupancyCount() int {
	if s.slotList == nil {
		return 0
	}

	return s.slotList.CountSelf()
}

// NewStorey returns a Storey object
func NewStorey(maxSlots int) *Storey {
	return &Storey{
		maxSlots: maxSlots,
	}
}

// Slot - each Storey has slots <= maxSlots
type Slot struct {
	prevSlot *Slot
	car      *Car
	position int
	nextSlot *Slot
}

// Leave - leave the Car, and connect the prev with next
func (s *Slot) Leave() error {
	return nil
}

// AddNext - add a new Slot after the current and associate the current next to the new.
func (s *Slot) AddNext(sc *Slot) error {
	return nil
}

// CountSelf counts 1 for self and relayes the count Self to next Slot.
func (s Slot) CountSelf() int {
	if s.nextSlot == nil {
		return 1
	}

	return 1 + s.nextSlot.CountSelf()
}

// NewSlot returns a slot object
func NewSlot(car *Car, position int) *Slot {
	return &Slot{car: car, position: position}
}

// Car - define the car preoperties
type Car struct {
	numberPlate string
	color       string
}

// NewCar returns a new car object
func NewCar(numberPlate, color string) *Car {
	return &Car{
		numberPlate: numberPlate,
		color:       color,
	}
}
