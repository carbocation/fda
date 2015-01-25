package fda

// DrugService provides data on drugs and drug safety event reporting
type DrugService struct {
	client *Client
}

// EventSearch finds events that have been reported while a patient is on medications.
// It yields a slice of SafetyReport, which is a structured response providing information
// about the event.
func (s *DrugService) EventSearch(search string, limit, skip int) ([]SafetyReport, *Meta, error) {
	data := []SafetyReport{}

	meta, err := s.client.search("/drug/event", search, limit, skip, &data)
	return data, meta, err
}

// LabelSearch finds drug labels, which is manufacturer-created information
func (s *DrugService) LabelSearch(search string, limit, skip int) ([]DrugLabel, *Meta, error) {
	data := []DrugLabel{}

	meta, err := s.client.search("/drug/label", search, limit, skip, &data)
	return data, meta, err
}

// EnforcementSearch finds reports of drug enforcement actions
func (s *DrugService) EnforcementSearch(search string, limit, skip int) ([]EnforcementReport, *Meta, error) {
	data := []EnforcementReport{}

	meta, err := s.client.search("/drug/enforcement", search, limit, skip, &data)
	return data, meta, err
}
