package simple

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

const baseURL = "http://api.wolframalpha.com/v1/simple"

// https://products.wolframalpha.com/simple-api/documentation/
type RestClient struct {
	client *http.Client

	AppID string

	Width    int    `json:"width"`
	Fontsize int    `json:"fontsize"`
	Units    string `json:"units"`
	Timeout  int    `json:"timeout"`
}

func New(appID string) *RestClient {
	return &RestClient{
		client:   &http.Client{},
		AppID:    appID,
		Width:    400,
		Fontsize: 14,
		Units:    "metric",
		Timeout:  5,
	}
}

func (r *RestClient) QueryParameters() url.Values {
	params := url.Values{}

	if r.AppID != "" {
		params.Add("appid", r.AppID)
	}

	if r.Width != 0 {
		params.Add("width", fmt.Sprintf("%d", r.Width))
	}

	if r.Fontsize != 0 {
		params.Add("fontsize", "16")
	}

	return params
}

func (r *RestClient) NewRequest(input string) (*http.Request, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		panic(err)
	}

	params := r.QueryParameters()
	params.Add("i", input)

	u.RawQuery = params.Encode()

	return http.NewRequest("GET", u.String(), nil)
}

func (w *RestClient) QueryToFile(input string, filename string) error {
	req, err := w.NewRequest(input)
	if err != nil {
		return err
	}

	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
