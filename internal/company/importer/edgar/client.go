package importer

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	endpoint *url.URL
	client   *http.Client
}

func NewEDGARClient(endpoint *url.URL) *Client {
	client := &http.Client{
		Timeout: 1 * time.Minute,
	}

	return &Client{
		endpoint: endpoint,
		client:   client,
	}
}

func (c *Client) GetBulkData(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.endpoint.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
