package fda

// DrugService provides data on drugs and drug safety event reporting
type DrugService struct {
	client *Client
}

// EventSearch finds events that have been reported while a patient is on medications.
// It yields a slice of DrugSafetyReport, which is a structured response providing information
// about the event.
func (s *DrugService) EventSearch(search string, limit, skip int) ([]DrugSafetyReport, *Meta, error) {
	data := []DrugSafetyReport{}

	meta, err := s.client.search("/drug/event", search, limit, skip, &data)
	return data, meta, err
}

// EventCount counts events that have been reported while a patient is on medications.
// Because the structure of the data depends on whether you are counting a field or a date,
// It expects you to formulate the correct structure and pass it as *data to be modified.
// Limit is permitted but Skip is not, and is therefore elided
func (s *DrugService) EventCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/drug/event", search, count, limit, &data)
	return meta, err
}

// LabelSearch finds drug labels, which is manufacturer-created information
func (s *DrugService) LabelSearch(search string, limit, skip int) ([]DrugLabel, *Meta, error) {
	data := []DrugLabel{}

	meta, err := s.client.search("/drug/label", search, limit, skip, &data)
	return data, meta, err
}

// LabelCount counts drug labels based on parameters you pass
// It expects you to formulate the correct structure and pass it as *data to be modified.
func (s *DrugService) LabelCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/drug/label", search, count, limit, &data)
	return meta, err
}

// EnforcementSearch finds reports of drug enforcement actions
func (s *DrugService) EnforcementSearch(search string, limit, skip int) ([]EnforcementReport, *Meta, error) {
	data := []EnforcementReport{}

	meta, err := s.client.search("/drug/enforcement", search, limit, skip, &data)
	return data, meta, err
}

// EnforcementCount counts drug enforcement actions based on parameters you pass
// It expects you to formulate the correct structure and pass it as *data to be modified.
func (s *DrugService) EnforcementCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/drug/enforcement", search, count, limit, &data)
	return meta, err
}
