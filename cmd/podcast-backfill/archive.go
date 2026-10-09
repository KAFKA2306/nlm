package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type archiveClient struct {
	identifier string
	accessKey  string
	secretKey  string
	http       *http.Client
}

func (a *archiveClient) publicURL(filename string) string {
	return fmt.Sprintf("https://archive.org/download/%s/%s", url.PathEscape(a.identifier), url.PathEscape(filename))
}

func (a *archiveClient) exists(ctx context.Context, filename string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, a.publicURL(filename), nil)
	if err != nil {
		return false, err
	}
	resp, err := a.http.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return false, fmt.Errorf("HEAD returned %s", resp.Status)
}

func (a *archiveClient) upload(ctx context.Context, filename string, data []byte) error {
	u := fmt.Sprintf("https://s3.us.archive.org/%s/%s", url.PathEscape(a.identifier), url.PathEscape(filename))
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, u, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "LOW "+a.accessKey+":"+a.secretKey)
	req.Header.Set("Content-Type", "audio/mp4")
	req.ContentLength = int64(len(data))
	req.Header.Set("x-amz-auto-make-bucket", "1")
	resp, err := a.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return fmt.Errorf("PUT returned %s: %s", resp.Status, strings.TrimSpace(string(body)))
	}
	return nil
}

func (a *archiveClient) waitAvailable(filename string, attempts int, delay time.Duration) error {
	for i := 0; i < attempts; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		exists, err := a.exists(ctx, filename)
		cancel()
		if err == nil && exists {
			return nil
		}
		if i+1 < attempts {
			time.Sleep(delay)
		}
	}
	return fmt.Errorf("Archive.org file did not become publicly available: %s", filename)
}
