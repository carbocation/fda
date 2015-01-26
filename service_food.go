package fda

// FoodService provides data on foods and food safety event reporting. Essentially
// it shows information on food recalls.
type FoodService struct {
	client *Client
}

// EnforcementSearch finds reports of food enforcement actions
func (s *FoodService) EnforcementSearch(search string, limit, skip int) ([]EnforcementReport, *Meta, error) {
	data := []EnforcementReport{}

	meta, err := s.client.search("/food/enforcement", search, limit, skip, &data)
	return data, meta, err
}

// EnforcementCount counts food enforcement actions based on parameters you pass
// It expects you to formulate the correct structure and pass it as *data to be modified.
func (s *FoodService) EnforcementCount(search, count string, data interface{}, limit int) (*Meta, error) {
	meta, err := s.client.count("/food/enforcement", search, count, limit, &data)
	return meta, err
}
