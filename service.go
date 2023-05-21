package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const contentType = "application/vnd.api + json"

type Client struct {
	BaseURL    string
	HttpClient http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HttpClient: http.Client{},
	}
}

func (c *Client) Create(acc AccountData) error {
	var data ResponseNotification
	data.Data = acc

	var bodyBuf bytes.Buffer
	err := json.NewEncoder(&bodyBuf).Encode(data)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/v1/organisation/accounts", c.BaseURL)
	resp, err := c.HttpClient.Post(url, contentType, &bodyBuf)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}

func (c *Client) Fetch(id string) (AccountData, error) {
	url := fmt.Sprintf("%s/v1/organisation/accounts/{%s}", c.BaseURL, id)

	resp, err := c.HttpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var responseNotification ResponseNotification
	err = json.NewDecoder(resp.Body).Decode(&responseNotification)
	if err != nil {
		return AccountData{}, err
	}
	return responseNotification.Data, nil
}

func (c *Client) Delete(id string, version string) error {
	url := fmt.Sprintf("%s/v1/organisation/accounts/{%s}?version={%s}", c.BaseURL, id, version)
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.HttpClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var responseNotification ResponseNotification
	err = json.NewDecoder(resp.Body).Decode(&responseNotification)
	if err != nil {
		return err
	}
	return nil
}
