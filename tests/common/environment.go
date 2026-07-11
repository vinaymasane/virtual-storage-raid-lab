package common

import (
	"os"
	"testing"
)

func RequireEnv(t *testing.T, key string) string {

	t.Helper()

	v := os.Getenv(key)

	if v == "" {
		t.Fatalf("%s environment variable not set", key)
	}

	return v
}