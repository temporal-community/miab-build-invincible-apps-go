package iplocate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetIP fetches the public IP address.
func GetIP(ctx context.Context) (string, error) {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ip := strings.TrimSpace(string(body))
	return ip, nil
}

// GetLocationInfo uses the IP address to fetch location information.
func GetLocationInfo(ctx context.Context, ip string) (string, error) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data struct {
		City       string `json:"city"`
		RegionName string `json:"regionName"`
		Country    string `json:"country"`
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s, %s, %s", data.City, data.RegionName, data.Country), nil
}
