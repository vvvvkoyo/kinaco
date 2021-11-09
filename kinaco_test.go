package kinaco

import (
	"strings"
	"testing"
)

func TestSuccess(t *testing.T) {
	result := Generate()
	if result == "" {
		t.Fatal("Expected not empty.")
	}

	parts := strings.Split(result, "-")
	if len(parts) != 3 {
		t.Fatal("Expected 3 parts.")
	}
}

func TestDelimiterSuccess(t *testing.T) {
	result := Delimiter("/").Generate()
	if result == "" {
		t.Fatal("Expected not empty.")
	}

	hyphenParts := strings.Split(result, "-")
	if len(hyphenParts) != 1 {
		t.Fatal("Expected 1 part.")
	}

	slashParts := strings.Split(result, "/")
	if len(slashParts) != 3 {
		t.Fatal("Expected 3 parts.")
	}
}

func TestNumSuccess(t *testing.T) {
	result := Num(5).Generate()
	if result == "" {
		t.Fatal("Expected not empty.")
	}

	parts := strings.Split(result, "-")
	if len(parts) != 5 {
		t.Fatal("Expected 5 part.")
	}

	// if num is 0, the result should be 3-3-3
	result2 := Num(0).Generate()
	if result2 == "" {
		t.Fatal("Expected not empty.")
	}

	parts2 := strings.Split(result2, "-")
	if len(parts2) != 3 {
		t.Fatal("Expected 3 part.")
	}
}

func TestNumDelimiterSuccess(t *testing.T) {
	result := Num(5).Delimiter("/").Generate()
	if result == "" {
		t.Fatal("Expected not empty.")
	}

	hyphenParts := strings.Split(result, "-")
	if len(hyphenParts) != 1 {
		t.Fatal("Expected 1 part.")
	}

	slashParts := strings.Split(result, "/")
	if len(slashParts) != 5 {
		t.Fatal("Expected 5 parts.")
	}
}
