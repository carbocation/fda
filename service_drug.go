package fda

import (
	"fmt"
	"net/url"
	"strconv"
)

type DrugService struct {
	client *Client
}

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
