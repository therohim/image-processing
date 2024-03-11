package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v -timeout 30s -run ^TestResizeImage$ Ubersnap-middle-backend-programmer-test/utils

func TestResizeImage(t *testing.T) {
	image1 := "../images/test.png"

	result1 := ResizeImage(image1, "../images")

	assert.Equal(t, "../images/test.jpeg", *result1)
}
