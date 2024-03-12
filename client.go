package tiktokbiz

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	utils "github.com/hung12ct/go-tiktok-business-sdk/internal"
)

// Client sdk client
type Client struct {
	config ClientConfig

	requestBuilder    utils.RequestBuilder
	createFormBuilder func(io.Writer) utils.FormBuilder
}

// NewClient creates new Tiktok Business API client.
func NewClient(appID, secret string) *Client {
	config := DefaultConfig(appID, secret)
	return NewClientWithConfig(config)
}

// NewClientWithConfig creates new Tiktok Business API for specified config.
func NewClientWithConfig(config ClientConfig) *Client {
	return &Client{
		config:         config,
		requestBuilder: utils.NewRequestBuilder(),
		createFormBuilder: func(body io.Writer) utils.FormBuilder {
			return utils.NewFormBuilder(body)
		},
	}
}

// set authentication token.
func (c *Client) SetAuthToken(authToken string) {
	c.config.AuthToken = authToken
}

// set base Url.
func (c *Client) SetBaseUrl(baseUrl string) {
	c.config.BaseURL = baseUrl
}

type requestOptions struct {
	body   any
	header http.Header
}

type requestOption func(*requestOptions)

func withBody(body any) requestOption {
	return func(args *requestOptions) {
		args.body = body
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, setters ...requestOption) (*http.Request, error) {
	// Default Options
	args := &requestOptions{
		body:   nil,
		header: make(http.Header),
	}
	for _, setter := range setters {
		setter(args)
	}
	req, err := c.requestBuilder.Build(ctx, method, url, args.body, args.header)
	if err != nil {
		return nil, err
	}
	c.setCommonHeaders(req)
	return req, nil
}

func (c *Client) sendRequest(req *http.Request) (*BaseResponseWithData, error) {
	req.Header.Set("Accept", "application/json")

	// Check whether Content-Type is already set, Upload Files API requires
	// Content-Type == multipart/form-data
	contentType := req.Header.Get("Content-Type")
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if isFailureStatusCode(res) {
		return nil, c.handleErrorResp(res)
	}

	var resp BaseResponseWithData
	err = decodeResponse(res.Body, &resp)
	if resp.IsError() {
		return nil, resp.BaseResponse
	}
	return &resp, err
}

func (c *Client) setCommonHeaders(req *http.Request) {
	if c.config.AuthToken != "" {
		req.Header.Set("Access-Token", c.config.AuthToken)
	}
}

func isFailureStatusCode(resp *http.Response) bool {
	return resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest
}

func decodeResponse(body io.Reader, v *BaseResponseWithData) error {
	return json.NewDecoder(body).Decode(v)
}

func unmarshalResponseData(resp json.RawMessage, respData any) error {
	if err := json.Unmarshal(resp, respData); err != nil {
		return err
	}
	return nil
}

func (c *Client) sendAndUnmarshal(req *http.Request, respData any) error {
	resp, err := c.sendRequest(req)
	if err != nil {
		return err
	}

	if err := unmarshalResponseData(resp.Data, respData); err != nil {
		return err
	}

	return nil
}

// fullURL returns full URL for request.
func (c *Client) fullURL(endpoint string) string {
	return fmt.Sprintf("%s/%s/%s", c.config.BaseURL, c.config.APIVersion, endpoint)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	var errRes ErrorResponse
	err := json.NewDecoder(resp.Body).Decode(&errRes)
	if err != nil || errRes.Error == nil {
		reqErr := &RequestError{
			HTTPStatusCode: resp.StatusCode,
			Err:            err,
		}
		if errRes.Error != nil {
			reqErr.Err = errRes.Error
		}
		return reqErr
	}

	errRes.Error.HTTPStatusCode = resp.StatusCode
	return errRes.Error
}
