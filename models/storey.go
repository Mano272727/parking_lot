package models

type Storey struct {
	maxSlots int
	slots    *Slot
}

// Park - check if the Slot is available
// if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Park(numberPlate, color string) (*Slot, error) {
	return &Slot{}, nil
}

// Park - check if the Slot is available
// if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Leave(numberPlate string) (*Slot, error) {
	return &Slot{}, nil
}

func (s *Storey) OccupancyCount() int {
	return 0
}

// NewStorey returns a Stirey object
func NewStorey(maxSlots int) *Storey {
	return &Storey{
		maxSlots: maxSlots,
	}
}

// Slot - each Storey has slots <= maxSlots
type Slot struct {
	prevSlot *Slot
	occupant Car
	id       int
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

// Car - define the car preoperties
type Car struct {
	numberPlate string
	color       string
}

func NewCar(numberPlate, color string) *Car {
	return &Car{
		numberPlate: numberPlate,
		color:       color,
	}
}
