package otpravka

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseUrl = "https://otpravka-api.pochta.ru/"

	ContentTypeJSON     = "application/json"
	ContentTypeXML      = "application/xml"
	ContentTypeTextHtml = "text/html"
)

type service struct {
	client *Client
}

type Client struct {
	client *http.Client

	BaseURL   *url.URL
	AuthToken string
	AuthKey   string

	common service // Reuse a single struct instead of allocating one for each service on the heap.
	Order  *OrderService
}

func NewClient(httpClient *http.Client, apiKey, apiToken string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(defaultBaseUrl)

	c := &Client{client: httpClient, BaseURL: baseURL, AuthKey: apiKey, AuthToken: apiToken}

	c.common.client = c
	c.Order = (*OrderService)(&c.common)

	return c
}

func (c *Client) NewRequest(method, urlStr, contentType string, body interface{}) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		var err error
		buf = &bytes.Buffer{}
		if hasContentType(contentType, ContentTypeJSON) {
			enc := json.NewEncoder(buf)
			err = enc.Encode(body)
		} else if hasContentType(contentType, ContentTypeXML) {
			enc := xml.NewEncoder(buf)
			err = enc.Encode(body)
		} else {
			return nil, fmt.Errorf("request Content-Type \"%q\" is unknown", contentType)
		}

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", contentType)
	}

	req.Header.Add("X-User-Authorization", "Basic "+c.AuthKey)
	req.Header.Add("Authorization", "AccessToken "+c.AuthToken)

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	var err error

	req = req.WithContext(ctx)

	response, err := c.client.Do(req)
	if response != nil {
		defer func() {
			if e := response.Body.Close(); e != nil && err == nil {
				err = e // if body not close, return err
			}
		}()
	}

	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	if v != nil {
		contentType := response.Header.Get("Content-Type")
		body, err0 := ioutil.ReadAll(response.Body)
		if err0 != nil {
			return response, err0
		}

		err = decodeBody(contentType, body, v)
	}

	return response, err
}

func decodeBody(contentType string, body []byte, v interface{}) error {
	var err error

	if hasContentType(contentType, ContentTypeJSON) {
		err = json.Unmarshal(body, v)
	} else if hasContentType(contentType, ContentTypeXML) {
		err = xml.Unmarshal(body, v)
	} else if hasContentType(contentType, ContentTypeTextHtml) {
		err = errors.New(string(body))
	} else {
		err = fmt.Errorf("response Content-Type %q is unknown", contentType)
	}

	if errors.Is(err, io.EOF) {
		err = nil // ignore EOF errors caused by empty response body
	} else if err != nil {
		e := &ErrorResponse{}
		err = decodeBody(contentType, body, e)
		if e.Code != "" {
			return e
		}
	}

	return err
}

func hasContentType(contentType, mimetype string) bool {
	for _, v := range strings.Split(contentType, ";") {
		t, _, err := mime.ParseMediaType(v)
		if err != nil {
			break
		}
		if t == mimetype {
			return true
		}
	}
	return false
}

type Error struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func (e Error) String() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Text)
}

func (e Error) Error() string {
	return e.String()
}

type ErrorResponse struct {
	Code    string `json:"code"`
	SubCode string `json:"sub-code"`
	Desc    string `json:"desc"`
}

func (e ErrorResponse) String() string {
	return fmt.Sprintf("[%s](%s) \"%s\"", e.Code, e.SubCode, e.Desc)
}

func (e ErrorResponse) Error() string {
	return e.String()
}
