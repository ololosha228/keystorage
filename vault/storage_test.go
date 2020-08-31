package vault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIt(t *testing.T) {
	c, err := New(
		"http://localhost:8200",
		"s.p0L1pGVlZtfm94evqktefCqQ",
		"/v1/secret/users")
	assert.NoError(t, err)

	c.HasService("")
}
