package strutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, IsEmpty(""), true)
	assert.Equal(t, IsEmpty("a"), false)
}

func TestChineseCount(t *testing.T) {
	assert.Equal(t, ChineseCount("我爱中国"), 4)
	assert.Equal(t, ChineseCount(" "), 0)
	assert.Equal(t, ChineseCount("我Aa'"), 1)
}

func TestIsContainChineseChar(t *testing.T) {
	assert.True(t, IsContainChineseChar("我爱中国"))
	assert.False(t, IsContainChineseChar("a"))
}

func TestGetFirstChineseChar(t *testing.T) {
	assert.Equal(t, GetFirstChineseChar("我爱中国"), "我")
	assert.Equal(t, GetFirstChineseChar("ABC我爱中国"), "我")
}

func TestGetChineseChar(t *testing.T) {
	assert.Equal(t, GetChineseChar("我爱中国"), []string{"我", "爱", "中", "国"})
}

func TestGetChineseString(t *testing.T) {
	assert.Equal(t, GetChineseString("Abc我爱De中国123"), "我爱中国")
}

func TestNoCaseEq(t *testing.T) {
	assert.True(t, NoCaseEq("abc", "Abc"))
	assert.False(t, NoCaseEq("abc", "Abv"))
}
