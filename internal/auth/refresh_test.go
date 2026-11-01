package auth

import (
	"testing"
)

func TestGenerateSAPISIDHASH(t *testing.T) {
	tests := []struct {
		name      string
		sapisid   string
		timestamp int64
		want      string
	}{
		{
			name:      "Current notebook.google.com origin",
			sapisid:   "ehxTF4-jACAOIp6k/Ax2l7oysalHiZneAB",
			timestamp: 1757337921,
			want:      "30fc826e39451a7a4cd75a0621013cc4afe1d6de",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &RefreshClient{
				sapisid: tt.sapisid,
			}

			got := client.generateSAPISIDHASH(tt.timestamp)
			if got != tt.want {
				t.Errorf("generateSAPISIDHASH() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCookieValue(t *testing.T) {
	cookies := "HSID=fake_hsid_value; SSID=fake_ssid_value; SAPISID=fake_sapisid_value; OTHER=value"

	tests := []struct {
		name   string
		cookie string
		want   string
	}{
		{"Extract SAPISID", "SAPISID", "fake_sapisid_value"},
		{"Extract HSID", "HSID", "fake_hsid_value"},
		{"Extract SSID", "SSID", "fake_ssid_value"},
		{"Non-existent cookie", "NOTFOUND", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractCookieValue(cookies, tt.cookie)
			if got != tt.want {
				t.Errorf("extractCookieValue(%s) = %v, want %v", tt.cookie, got, tt.want)
			}
		})
	}
}
