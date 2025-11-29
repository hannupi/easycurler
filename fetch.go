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
			panic(fmt.Sprintf("Error creating request: %s", err))
		}

		client := &http.Client{Timeout: 30 * time.Second}
		res, err := client.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Error creating request: %s", err))
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(fmt.Sprintf("Unable to read res body: %s", err))
		}

		if res.StatusCode != http.StatusOK {
			return httpResMsg(fmt.Sprintf("HTTP error: %s", res.Status))
		}
		return httpResMsg(strings.TrimSpace(string(body)))
	}
}
