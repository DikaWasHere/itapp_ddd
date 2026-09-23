package sso

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
 
	domainApp "itapp_ddd/domain/application"
 
	"github.com/joho/godotenv"
)
 
type Client struct {
	httpClient *http.Client
	url        string
}
 
func NewClient() *Client {
	_ = godotenv.Load()
	return &Client{
		httpClient: &http.Client{Timeout: 15 * time.Second},
		url:        os.Getenv("URL_SSO"),
	}
}
 
type ssoResponse struct {
	Status  int64     `json:"status"`
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    []ssoData `json:"data"`
}
 
type ssoData struct {
	ID     uint   `json:"id"`
	Module string `json:"module"`
}
 
// FetchApplications memanggil API SSO dan mengembalikan daftar aplikasi
func (c *Client) FetchApplications(ctx context.Context, req domainApp.GetListApplicationRequest) ([]domainApp.SSOApplication, error) {
	if c.url == "" {
		return nil, fmt.Errorf("URL_SSO is not set in .env")
	}
 
	payload, err := json.Marshal(map[string]interface{}{
		"username":      req.Username,
		"session_token": req.SessionToken,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal payload: %w", err)
	}
 
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
 
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("key-puninar", "73b3b0d9edd2de19566f32d6e5a3be1047125300d808f35286276a0f56e69be4")
	httpReq.Header.Set("timestamp", "2023-09-19 15:44:21")
 
	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("call SSO: %w", err)
	}
	defer resp.Body.Close()
 
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
 
	var parsed ssoResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}
 
	out := make([]domainApp.SSOApplication, 0, len(parsed.Data))
	for _, d := range parsed.Data {
		out = append(out, domainApp.SSOApplication{ID: d.ID, Module: d.Module})
	}
	return out, nil
}
