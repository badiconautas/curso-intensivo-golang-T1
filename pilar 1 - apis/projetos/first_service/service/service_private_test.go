package service

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
)

func TestValidate(t *testing.T) {
	validNotebookName := "Caderno de estudo"
	err := validateNotebookName(validNotebookName)
	if err != nil {
		t.Fatalf("Notebook name deveria ser valido %v", err)
	}
}