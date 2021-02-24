package resources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOk(t *testing.T) {
	fileName := "../app.yml"
	got := YmlDecode(fileName)
	want := "Martin D'vloper"

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestWrongFileName(t *testing.T) {
	fileName := "../apps.yml"
	assert.Panics(t, func() {
		YmlDecode(fileName)
	})
}
