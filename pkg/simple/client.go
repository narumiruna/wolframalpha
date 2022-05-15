package simple

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseURL = "http://api.wolframalpha.com/v1/simple"

// https://products.wolframalpha.com/simple-api/documentation/
type RestClient struct {
	client *http.Client

	AppID string
}

func New(appID string) *RestClient {
	return &RestClient{
		client: &http.Client{},
		AppID:  appID,
	}
}

func (r *RestClient) GetParameters(input string, options *QueryOptions) url.Values {
	params := url.Values{}

	if input != "" {
		params.Add("i", input)
	}

	if r.AppID != "" {
		params.Add("appid", r.AppID)
	}

	if options == nil {
		return params
	}

	if options.Width != 0 {
		params.Add("width", fmt.Sprintf("%d", options.Width))
	}

	if options.Fontsize != 0 {
		params.Add("fontsize", fmt.Sprintf("%d", options.Fontsize))
	}

	if options.Units != "" {
		params.Add("units", options.Units)
	}

	if options.Timeout != 0 {
		params.Add("timeout", fmt.Sprintf("%d", options.Timeout))
	}

	return params
}

func (r *RestClient) NewRequest(params url.Values) (*http.Request, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	if params != nil {
		u.RawQuery = params.Encode()
	}

	return http.NewRequest("GET", u.String(), nil)
}

func (c *RestClient) Query(input string, options *QueryOptions) ([]byte, error) {
	params := c.GetParameters(input, options)

	req, err := c.NewRequest(params)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (c *RestClient) QueryFile(input string, filename string, options *QueryOptions) error {
	body, err := c.Query(input, options)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, body, 0644)
}
