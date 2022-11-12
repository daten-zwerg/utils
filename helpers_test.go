package helpers

import (
	"os"
	"testing"
)

func TestNewMongodbUri(t *testing.T) {
	os.Setenv("MONGODB_PWD", "1234")
	os.Setenv("MONGODB_NAME", "test")
	want := "mongodb://localhost:27017/test"

	got := NewMongodbUri()

	if got != want {
		t.Errorf("invalid created uri, got %s, want %s", got, want)
	}
}
