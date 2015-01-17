package fda

/*
import (
	"encoding/xml"
	"fmt"
	"net/url"
)

type SearchService struct {
	client *Client
}

type SearchData struct {
	XMLName xml.Name `xml:"searchlist"`
	MRN     string   `xml:"mrn"`
	Site    string   `xml:"site"`
	Name    string   `xml:"name"`
	Request string   `xml:"request"`
	Query   []Query  `xml:"query"`
}

type Query struct {
	Search     string       `xml:"search"`
	SearchItem []SearchItem `xml:"searchitem"`
}

type SearchItem struct {
	Type         string `xml:"type"`
	MID          string
	ScheduleDate string `xml:"Scheduledate"`
	Site         string `xml:"site"`
	Status       string `xml:"status"`
	Desc         string `xml:"desc"`
	Context      string `xml:"context"`
	Highlight    string `xml:"highlight"`
}

func (s *SearchService) Search(query string, mrn MRN) (SearchData, error) {
	params := url.Values{}
	params.Add("function", "search")
	params.Add("MRN", mrn.Number)
	params.Add("site", mrn.Site)
	params.Add("search", query)

	u := "?" + params.Encode()

	req, err := s.client.NewRequest(u, "")
	if err != nil {
		return SearchData{}, err
	}

	data := SearchData{}
	_, err = s.client.Do(req, &data)
	if err != nil {
		return SearchData{}, fmt.Errorf("SearchService:Search err: %s", err)
	}

	return data, nil
}
*/
