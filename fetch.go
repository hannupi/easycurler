package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"io"
	"net/http"
	"strings"
)

func fetchURL(url string) tea.Cmd {
	return func() tea.Msg {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			return errMsg(err)
		}
		body, err := io.ReadAll(res.Body)

		if res.StatusCode != http.StatusOK {
			return errMsg(fmt.Errorf("HTTP error: %s", res.Status))
		}
		return httpResMsg(strings.TrimSpace(string(body)))
	}
}
