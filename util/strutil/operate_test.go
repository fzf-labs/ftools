package strutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuote(t *testing.T) {
	assert.Equal(t, Quote("abc"), "\"abc\"")
}

func TestAddSlashes(t *testing.T) {
	assert.Equal(t, AddSlashes(`{"key": 123}`), `{\"key\": 123}`)
}

func TestStripSlashes(t *testing.T) {
	assert.Equal(t, StripSlashes(`{\"key\": 123}`), `{"key": 123}`)
}

func TestUpperEnglishWord(t *testing.T) {
	assert.Equal(t, UpperEnglishWord(`wo...shi...中文`), `Wo...Shi...中文`)
}

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, SnakeCase("rangePrice"), "range_price")
}
