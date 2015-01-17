/*
Please see https://open.fda.gov/api/reference/#query-syntax for how
queries should be structured
*/
package fda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	LibraryVersion = "0.0.1"

	// This is the FDA API endpoint
	BaseURL = "https://api.fda.gov"

	UserAgent = "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0)"
)

type Client struct {
	// HTTP client used to communicate with the API
	client *http.Client

	// Base URL for API requests
	BaseURL *url.URL

	// User Agent used when communicating with the API
	UserAgent string

	// Until Partners implements OAuth
	APIKey string

	// Accept responses of this type. Defaults to JSON. Currently FDA only supports JSON.
	ResponseType string

	// For accessing functions of the API
	Drug   *DrugService
	Device *DeviceService
	Food   *FoodService

	// Temporary Response
	Response *Response
}

// NewClient returns a new API client. if a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, apikey string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: UserAgent,
		APIKey:    apikey,
	}

	c.Drug = &DrugService{client: c}
	c.Device = &DeviceService{client: c}
	c.Food = &FoodService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in endpoint,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(endpoint string, params url.Values, body string) (*http.Request, error) {
	if c.APIKey != "" {
		params.Add("api_key", c.APIKey)
	}

	rel, err := url.Parse(endpoint + "?" + params.Encode())
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)
	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	//For debugging
	//buffer := &bytes.Buffer{}
	//buffer.ReadFrom(resp.Body)
	//fmt.Println(buffer.String())

	r := &Response{HTTPResponse: resp}
	if v != nil {
		// Our passed argument is now the value of r.Response
		r.Results = v
		// Decode the JSON into r.Data (which is our argument)
		err = json.NewDecoder(resp.Body).Decode(r)
		c.Response = r
	}
	return resp, err
}

// Reponse resembles the returned packets
type Response struct {
	HTTPResponse *http.Response
	Response     interface{}
	Meta         *Meta
	Results      interface{}
}

// Meta represents information about the response. If all goes well,
// only a Code key with value 200 will present. However, sometimes things
// go wrong, and in that case ErrorType and ErrorMessage are present.
type Meta struct {
	Disclaimer  string
	License     string
	LastUpdated string `json:"last_updated"`
	Results     struct {
		Skip  int
		Limit int
		Total int
	}
}

// ErrorResponse represents a Response which contains an error
type ErrorResponse struct {
	HTTPResponse *http.Response
	Err          struct {
		StatusCode int
		Code       string
		Message    string
	} `json:"error"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %v",
		r.HTTPResponse.Request.Method, r.HTTPResponse.Request.URL.String(),
		r.Err.StatusCode, r.Err.Code, r.Err.Message)
}

// CheckResponse checks the API response for error, and returns it
// if present. A response is considered an error if it has non StatusOK
// code.
func CheckResponse(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}

	resp := new(ErrorResponse)
	resp.HTTPResponse = r

	err := json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return fmt.Errorf("CheckResponse:error: %s", err.Error())
	}

	resp.Err.StatusCode = r.StatusCode

	return resp
}
