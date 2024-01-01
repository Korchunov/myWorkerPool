package coincap

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type client struct {
	c *http.Client
}

func newClient(timeout time.Duration) (*client, error) {
	if timeout == 0 {
		return nil, errors.New("УПС")
	}
	return &client{
		c: &http.Client{
			Timeout: timeout,
		},
	}, nil
}

func (client *client) GetAssets() Assets {
	baseUrl := "https://api.coincap.io/v2/assets"
	req, err := http.NewRequest("GET", baseUrl, bytes.NewBuffer(nil))
	if err != nil {
		log.Fatal("УПС 1")
	}
	resp, err := client.c.Do(req)
	if err != nil {
		log.Fatal("УПС 2")
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Увы и ах")
	}
	var a Assets
	json.Unmarshal(respBody, &a)
	return a
}
