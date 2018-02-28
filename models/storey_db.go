package models

var (
	// RuleEvenDistribution - rule set to evenly distribute the cars accross parking lot.
	RuleEvenDistribution = "even_distribution"
)

// storeyDB holds the data in memory while run time.
// in requirements document multi storey is mentioned.
// but no further actions are requested in the same.
// so the Storey is defined as an array.
type storeyDB struct {
	Storeys []*Storey
	Rule    string
}

// NewStoreyRunTimeDB returns an instance of the storey db.
func NewStoreyRunTimeDB(maxSlots int) *storeyDB {
	storey := NewStorey(maxSlots)
	db := &storeyDB{
		Storeys: []*Storey{
			storey,
		},
	}
	storey.db = db
	return db
}

func (s *storeyDB) AddStorey(maxSlots int) (StoreyResponse, error) {
	storey := NewStorey(maxSlots)
	storey.db = s
	s.Storeys = append(s.Storeys, storey)
	return StoreyResponse{
		slots:   []Slot{},
		storey:  storey,
		command: CmdCreateParkingLot,
	}, nil
}

// Park a car
func (s *storeyDB) Park(numberPlate, color string) (StoreyResponse, error) {
	// until we start supporting more than one storey
	slot, err := s.Storeys[0].Park(numberPlate, color)
	sResponse := StoreyResponse{
		slots: []Slot{
			*slot,
		},
		command: CmdPark,
	}

	return sResponse, err
}

// LeaveByPosition leave a car froma slot by the position
func (s *storeyDB) LeaveByPosition(position int) (StoreyResponse, error) {
	slot, err := s.Storeys[0].LeaveByPosition(position)
	sResponse := StoreyResponse{
		slots: []Slot{
			*slot,
		},
		command: CmdLeave,
	}

	return sResponse, err
}

// FindByRegistrationNumber find slot having car with registration number.
func (s *storeyDB) FindByRegistrationNumber(numberPlate string) (StoreyResponse, error) {
	slot, err := s.Storeys[0].FindByRegistrationNumber(numberPlate)
	sResponse := StoreyResponse{
		slots: []Slot{
			*slot,
		},
		command: CmdSlotnoByRegNumber,
	}

	return sResponse, err
}

// FindAllByColor find the slots having cars with the color.
func (s *storeyDB) FindAllByColor(color, cmd string) (StoreyResponse, error) {
	slots, err := s.Storeys[0].FindAllByColor(color)
	sResponse := StoreyResponse{
		slots:   slots,
		command: cmd,
	}

	return sResponse, err
}

// All returns the slots
func (s *storeyDB) All() (StoreyResponse, error) {
	slots, err := s.Storeys[0].AllSlots()
	return StoreyResponse{
		slots:   slots,
		command: CmdStatus,
	}, err
}
