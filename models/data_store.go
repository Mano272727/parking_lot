package ds

import "parking_lot/models"

type dataStore interface {
	Park(string, string) (*models.Slot, error)
	LeaveByPosition(int) (*models.Slot, error)
	FindByRegistrationNumber(string) (*models.Slot, error)
	FindAllByColor(string) ([]*models.Slot, error)
}
