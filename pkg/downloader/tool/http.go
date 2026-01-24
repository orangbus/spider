package tool

import (
	"crypto/tls"
	//"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
	//"time"
)

func Get(url string) (io.ReadCloser, error) {
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//if resp.StatusCode != 200 {
	//	return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	//}
	//return resp.Body, nil

	c := http.Client{
		Timeout: time.Duration(120) * time.Second,
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
