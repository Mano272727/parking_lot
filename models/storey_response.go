package models

type StoreyResponse struct {
	slots   []Slot
	command string
}

func (s *StoreyResponse) String() string {
	return ""
}
