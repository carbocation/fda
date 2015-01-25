package fda

import (
	"fmt"
	"net/url"
	"strconv"
)

// DrugService provides data on drugs and drug safety event reporting
type DrugService struct {
	client *Client
}

// EventSearch finds events that have been reported while a patient is on medications.
// It yields a slice of SafetyReport, which is a structured response providing information
// about the event. Subjectively, while there are many fields in each SafetyReport, often
// the majority of fields are left empty.
func (s *DrugService) EventSearch(search string, limit, skip int) ([]SafetyReport, *Meta, error) {
	data := []SafetyReport{}

	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("skip", strconv.Itoa(skip))
	if search != "" {
		params.Add("search", search)
	}

	req, err := s.client.NewRequest("/drug/event", params, "")
	if err != nil {
		return data, &Meta{}, err
	}

	resp, err := s.client.Do(req, &data)
	if err != nil {
		return data, &Meta{}, fmt.Errorf("EventSearch:err: %s", err)
	}

	return data, resp.Meta, nil
}
