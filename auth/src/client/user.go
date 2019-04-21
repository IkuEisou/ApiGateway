// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Authentication API": User Resource Client
//
// Command:
// $ goagen
// --design=github.com/ikueisou/apigateway/auth/design
// --force=true
// --out=$(GOPATH)/src/github.com/ikueisou/apigateway/auth/src
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// LoginUserPath computes a request path to the login action of User.
func LoginUserPath() string {

	return fmt.Sprintf("/api/login")
}

// Sign a user in
func (c *Client) LoginUser(ctx context.Context, path string, payload *LoginPayload) (*http.Response, error) {
	req, err := c.NewLoginUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUserRequest create the request corresponding to the login action endpoint of the User resource.
func (c *Client) NewLoginUserRequest(ctx context.Context, path string, payload *LoginPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// RegisterUserPath computes a request path to the register action of User.
func RegisterUserPath() string {

	return fmt.Sprintf("/api/register")
}

// Create a new user
func (c *Client) RegisterUser(ctx context.Context, path string, payload *RegisterPayload) (*http.Response, error) {
	req, err := c.NewRegisterUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRegisterUserRequest create the request corresponding to the register action endpoint of the User resource.
func (c *Client) NewRegisterUserRequest(ctx context.Context, path string, payload *RegisterPayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}
