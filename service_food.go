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
