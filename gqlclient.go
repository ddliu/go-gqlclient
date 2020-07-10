package gqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/ddliu/fractal"
)

func New(options Options) *Client {
	return &Client{
		options,
	}
}

type Options struct {
	Endpoint string
	Header   http.Header
	Timeout  time.Duration
}

type Client struct {
	Options Options
}

type GraphqlError struct {
	errorsWrap *fractal.Context
}

func (e *GraphqlError) Error() string {
	return e.errorsWrap.String("0.message")
}

func (c *Client) Query(ctx context.Context, query string, variables interface{}) (*fractal.Context, error) {
	reqData := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	reqBody, _ := json.Marshal(reqData)
	req, err := http.NewRequestWithContext(ctx, "POST", c.Options.Endpoint, bytes.NewBuffer([]byte(reqBody)))
	if err != nil {
		return nil, err
	}

	for k, v := range c.Options.Header {
		if len(v) > 0 {
			req.Header.Set(k, v[0])
		}
	}

	httpclient := &http.Client{Timeout: c.Options.Timeout}
	response, err := httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Request error, status code " + strconv.Itoa(response.StatusCode))
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	bodyWrap := fractal.FromJson(body)
	if bodyWrap.Exist("errors") {
		return bodyWrap.GetContext("data"), &GraphqlError{
			errorsWrap: bodyWrap.GetContext("errors"),
		}
	}

	return bodyWrap.GetContext("data"), nil
}
