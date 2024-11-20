package api

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/giuszeppe/github-activity-go-cli/config"
)

type API struct {
	url     string
	headers []string
}

func GetAPI() API {
	key, err := config.ReadAPIKey()
	if err != nil {
		panic("No API key")
	}
	return API{url: "https://api.github.com", headers: []string{"-H Accept: application/vnd.github+json", "-H Authorization: Bearer " + key.APIKey}}
}

func (api *API) Fetch(path string) ([]byte, error) {
	curl := exec.Command("curl", strings.Join(api.headers, " "), api.url+path)
	out, err := curl.Output()
	if err != nil {
		return []byte{}, fmt.Errorf("curl error")
	}

	return out, nil
}
