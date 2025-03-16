//go:build !darwin

package notify

// Notify sends a notification to the system.
// Not implemented for other platforms.
func Notify(title, message string, icon []byte) error {
	return nil
}
