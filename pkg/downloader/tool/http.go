package tool

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string) (io.ReadCloser, error) {
	c := http.Client{
		Timeout: time.Duration(60) * time.Second,
		// 忽略 https
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
