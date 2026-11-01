package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testCompleteAuthCookies() string {
	return "SAPISID=api; SID=sid; HSID=hsid; SSID=ssid; APISID=apisid; other=value"
}

func TestValidateAuthCookies(t *testing.T) {
	complete := testCompleteAuthCookies()
	if err := validateAuthCookies(complete); err != nil {
		t.Fatalf("complete cookies rejected: %v", err)
	}

	tests := []struct {
		name    string
		cookies string
		missing string
	}{
		{"no cookies", "", "SAPISID"},
		{"missing SAPISID", strings.Replace(complete, "SAPISID=api; ", "", 1), "SAPISID"},
		{"missing SID", strings.Replace(complete, "SID=sid; ", "", 1), "SID"},
		{"missing HSID", strings.Replace(complete, "HSID=hsid; ", "", 1), "HSID"},
		{"missing SSID", strings.Replace(complete, "SSID=ssid; ", "", 1), "SSID"},
		{"missing APISID", strings.Replace(complete, "APISID=apisid; ", "", 1), "APISID"},
		{"empty value", strings.Replace(complete, "SID=sid", "SID=", 1), "SID"},
		{"similar name", strings.Replace(complete, "SID=sid", "X_SID=sid", 1), "SID"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAuthCookies(tt.cookies)
			if err == nil || !strings.Contains(err.Error(), "missing "+tt.missing) && !strings.Contains(err.Error(), ", "+tt.missing) {
				t.Fatalf("expected missing %s error, got %v", tt.missing, err)
			}
			if !strings.Contains(err.Error(), "nlm auth login") {
				t.Fatalf("missing reauthentication guidance: %v", err)
			}
			if strings.Contains(err.Error(), "sid=ssid") {
				t.Fatalf("error should not expose cookie values: %v", err)
			}
		})
	}
}

func TestPersistAuthRejectsIncompleteWithoutOverwriting(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	envPath := filepath.Join(home, ".nlm", "env")
	if err := os.MkdirAll(filepath.Dir(envPath), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(envPath, []byte("previous credentials"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, _, err := persistAuthToDisk("SID=sid", "token", "", "", ""); err == nil {
		t.Fatal("incomplete cookies were persisted")
	}
	got, err := os.ReadFile(envPath)
	if err != nil || string(got) != "previous credentials" {
		t.Fatalf("existing credentials changed on failure: %v", err)
	}
}

func TestDetectAuthInfoRequiresCompleteCookies(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	bad := "curl 'https://notebook.google.com/?at=testtoken&foo=1' -H 'cookie: SID=sid; HSID=hsid'"
	if _, _, err := detectAuthInfo(bad); err == nil {
		t.Fatal("incomplete stdin cookies accepted")
	}
	if _, err := os.Stat(filepath.Join(home, ".nlm", "env")); !os.IsNotExist(err) {
		t.Fatalf("incomplete cookies created credentials file: %v", err)
	}

	good := "curl 'https://notebook.google.com/?at=testtoken&foo=1' -H 'cookie: " + testCompleteAuthCookies() + "'"
	token, cookies, err := detectAuthInfo(good)
	if err != nil || token != "testtoken" || cookies != testCompleteAuthCookies() {
		t.Fatalf("complete stdin credentials rejected: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(home, ".nlm", "env"))
	if err != nil || !strings.Contains(string(data), "NLM_AUTH_TOKEN=") {
		t.Fatalf("complete credentials were not stored: %v", err)
	}
}

func TestRefreshRejectsIncompleteCookiesBeforeNetwork(t *testing.T) {
	t.Setenv("HOME", t.TempDir())
	t.Setenv("NLM_COOKIES", "SID=sid")
	err := refreshCredentials(false)
	if err == nil || !strings.Contains(err.Error(), "nlm auth login") {
		t.Fatalf("refresh should request reauthentication: %v", err)
	}
}
