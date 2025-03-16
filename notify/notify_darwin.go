//go:build darwin

package notify

import (
	"fmt"
	"os/exec"
)

// Notify sends a notification to the system.
func Notify(title, message string) error {
	script := fmt.Sprintf(`display notification "%s" with title "%s"`, message, title)
	cmd := exec.Command("osascript", "-e", script)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to send notification: %w", err)
	}

	return nil
}
