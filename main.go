package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func splitPanes(parentPaneID string, args ...string) (string, error) {
	cmdArgs := []string{"cli", "split-pane", "--pane-id", parentPaneID}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("wezterm", cmdArgs...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	paneID := strings.TrimSpace(out.String())
	return paneID, nil
}

func sendCommandToPane(paneID string, command string, wg *sync.WaitGroup) error {
	defer wg.Done()
	// direnv が読み込まれるのを待つ
	time.Sleep(500 * time.Microsecond)
	cmd := exec.Command("wezterm", "cli", "send-text", "--pane-id", paneID, "--no-paste")
	cmd.Stdin = strings.NewReader(command + "\n")
	return cmd.Run()
}

func main() {
	p0 := os.Getenv("WEZTERM_PANE")
	if p0 == "" {
		//TODO: 色つける
		fmt.Fprintln(os.Stderr, "This tool requires WezTerm")
		os.Exit(1)
	}
	// ペインを分割
	p1, err := splitPanes(p0, "--right")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// もう一回ペインを分割
	p2, err := splitPanes(p1, "--bottom")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// 各ペインで Gemini CLI を実行
	panes := []string{p0, p1, p2}

	var wg sync.WaitGroup

	for _, paneID := range panes {
		wg.Add(1)
		go sendCommandToPane(paneID, "gemini", &wg)
	}
	wg.Wait()
}
