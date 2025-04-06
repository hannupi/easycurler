package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func fetchURL(url string, reqMethod string) tea.Cmd {
	method := reqMethod
	return func() tea.Msg {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return errMsg(err)
		}

		client := &http.Client{Timeout: 30 * time.Second}
		res, err := client.Do(req)
		if err != nil {
			return errMsg(err)
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)

		if res.StatusCode != http.StatusOK {
			return errMsg(fmt.Errorf("HTTP error: %s", res.Status))
		}
		return httpResMsg(strings.TrimSpace(string(body)))
	}
}
