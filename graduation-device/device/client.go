package device

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	IP         string
	Port       string
	Url        string
	Protocol   string
	Method     string
	HttpClient *http.Client
}

func NewDeviceClient(url string) *Client {
	return &Client{

		Url: url,
	}
}

func (c *Client) Send(data []byte) error {

	_, err := http.Post(c.Url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	return err
}

func (c *Client) getHttpClient() {
	c.HttpClient = &http.Client{}
}

func (c *Client) GenerateData(data interface{}) (byteData []byte, err error) {
	if data == nil {
		return nil, errors.New("data is null")
	}
	byteData, err = json.Marshal(data)
	fmt.Println(string(byteData))
	return
}
