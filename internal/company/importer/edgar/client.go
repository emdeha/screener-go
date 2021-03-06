package edgarimporter

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

	// userAgent is required by EDGAR so they can rate limit requests.
	userAgent string
}

func NewEDGARClient(endpoint *url.URL, userAgent string) *Client {
	client := &http.Client{
		// We're going to download a big file, so we must have gracious timeout.
		Timeout: 30 * time.Minute,
	}

	return &Client{
		endpoint:  endpoint,
		client:    client,
		userAgent: userAgent,
	}
}

func (c *Client) GetBulkData(ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", c.endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	c.addHeadersRequiredByEDGAR(req)

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

func (c *Client) addHeadersRequiredByEDGAR(req *http.Request) {
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Host", "www.sec.gov")
}
