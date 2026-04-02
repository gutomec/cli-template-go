package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloCmd(t *testing.T) {
	assert.NotNil(t, helloCmd)
	assert.Equal(t, "hello", helloCmd.Use)
}
