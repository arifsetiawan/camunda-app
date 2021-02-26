package camunda

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

// Client ...
type Client struct {
	APIURL string
	Client *resty.Client
}

// NewClient ...
func NewClient(url string) *Client {
	c := &Client{
		APIURL: url,
	}

	c.Client = resty.New()
	return c
}

// IdentityVerify ...
func (c *Client) IdentityVerify(param *IdentityVerifyRequest) (response *IdentityVerifyResponse, err error) {
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetResult(&IdentityVerifyResponse{}).
		Post(c.APIURL + "/identity/verify")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return resp.Result().(*IdentityVerifyResponse), err
}

// ListTask ...
func (c *Client) ListTask(param *ListTaskRequest) (response *[]UserTask, err error) {
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetResult(&[]UserTask{}).
		Post(c.APIURL + "/task")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return resp.Result().(*[]UserTask), err
}

// FetchAndLockExternalTask ...
func (c *Client) FetchAndLockExternalTask(param *FetchAndLockRequest) (response *[]ExternalTask, err error) {
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		SetResult(&[]ExternalTask{}).
		Post(c.APIURL + "/external-task/fetchAndLock")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}

	return resp.Result().(*[]ExternalTask), err
}

// CompleteExternalTask ...
func (c *Client) CompleteExternalTask(id string, param *CompleteExternalTaskRequest) (err error) {
	resp, err := c.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(param).
		Post(c.APIURL + "/external-task/" + id + "/complete")
	if err != nil {
		return err
	}

	if resp.StatusCode() != 204 {
		return errors.New("Complete failed with response " + resp.String())
	}

	return nil
}
