package validate

import (
	"path/filepath"
	"testing"
)

func TestValidateSpecFile(t *testing.T) {
	root := filepath.Join("..", "..")
	if err := ValidateSpecFile(root); err != nil {
		t.Fatal(err)
	}
}

func TestValidateVectors(t *testing.T) {
	if err := ValidateVectors(); err != nil {
		t.Fatal(err)
	}
}
