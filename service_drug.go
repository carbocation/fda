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
// about the event.
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

// LabelSearch finds drug labels, which is manufacturer-created information
func (s *DrugService) LabelSearch(search string, limit, skip int) ([]DrugLabel, *Meta, error) {
	data := []DrugLabel{}

	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("skip", strconv.Itoa(skip))
	if search != "" {
		params.Add("search", search)
	}

	req, err := s.client.NewRequest("/drug/label", params, "")
	if err != nil {
		return data, &Meta{}, err
	}

	resp, err := s.client.Do(req, &data)
	if err != nil {
		return data, &Meta{}, fmt.Errorf("LabelSearch:err: %s", err)
	}

	return data, resp.Meta, nil
}

// EnforcementSearch finds reports of drug enforcement actions
func (s *DrugService) EnforcementSearch(search string, limit, skip int) ([]EnforcementReport, *Meta, error) {
	data := []EnforcementReport{}

	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("skip", strconv.Itoa(skip))
	if search != "" {
		params.Add("search", search)
	}

	req, err := s.client.NewRequest("/drug/enforcement", params, "")
	if err != nil {
		return data, &Meta{}, err
	}

	resp, err := s.client.Do(req, &data)
	if err != nil {
		return data, &Meta{}, fmt.Errorf("EnforcementSearch:err: %s", err)
	}

	return data, resp.Meta, nil
}
