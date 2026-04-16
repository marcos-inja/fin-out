package validate

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindModuleRoot walks up from start until it finds go.mod.
func FindModuleRoot(start string) (string, error) {
	dir := start
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("go.mod not found from %s", start)
}
