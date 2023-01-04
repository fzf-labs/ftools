package urlutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlEncode(t *testing.T) {
	assert.Equal(t, UrlEncode("www.baidu.com/?query=golang"), "www.baidu.com/?query%3Dgolang")
}

func TestUrlDecode(t *testing.T) {
	assert.Equal(t, UrlDecode("www.baidu.com/?query%3Dgolang"), "www.baidu.com/?query=golang")
}
