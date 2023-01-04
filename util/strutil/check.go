package strutil

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsEmpty 是否是空字符串
func IsEmpty(s string) bool {
	if s == "" {
		return true
	}
	return strings.TrimSpace(s) == ""
}

// ChineseCount 中文字符统计
func ChineseCount(str string) int {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	return count
}

// IsContainChineseChar 是否包含中文字符
func IsContainChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return true
		}
	}
	return false
}

// GetFirstChineseChar 返回第一个中文字符
func GetFirstChineseChar(str string) string {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			return fmt.Sprintf("%c", r)
		}
	}
	return ""
}

// GetChineseChar 过滤返回中文字符切片
func GetChineseChar(str string) []string {
	ss := make([]string, 0)
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			ss = append(ss, fmt.Sprintf("%c", r))
		}
	}
	return ss
}

// GetChineseString 过滤返回中文字符
func GetChineseString(str string) string {
	var ss string
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || (regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(r))) {
			ss += fmt.Sprintf("%c", r)
		}
	}
	return ss
}

// NoCaseEq 检查两个字符串是否相等且不区分大小写
func NoCaseEq(s, t string) bool {
	return strings.EqualFold(s, t)
}

// IsNumChar 如果给定字符是数字，则返回 true，否则返回 false。
func IsNumChar(c byte) bool {
	return c >= '0' && c <= '9'
}

// IsNumeric 如果给定的字符串是数字，则返回 true，否则返回 false。
func IsNumeric(s string) bool {
	return regexp.MustCompile(`^\d+$`).MatchString(s)
}

// IsAlphabet 判断字节是 字母
func IsAlphabet(char uint8) bool {
	// A 65 -> Z 90
	if char >= 'A' && char <= 'Z' {
		return true
	}

	// a 97 -> z 122
	if char >= 'a' && char <= 'z' {
		return true
	}
	return false
}

// IsAlphaNum 判断字节是 ASCII 字母、数字还是下划线
func IsAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}

// HasOneSub 判断给定字符串是否有子字符串。
func HasOneSub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

// HasAllSubs 给定字符串中的所有子字符串。
func HasAllSubs(s string, subs []string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// HasOnePrefix 字符串以其中一个子项开头
func HasOnePrefix(s string, prefixes []string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// IsValidUtf8 有效的 utf8 字符串检查
func IsValidUtf8(s string) bool {
	return utf8.ValidString(s)
}

// RuneIsWord char: a-zA-Z
func RuneIsWord(c rune) bool {
	return RuneIsLower(c) || RuneIsUpper(c)
}

// RuneIsLower char
func RuneIsLower(c rune) bool {
	return 'a' <= c && c <= 'z'
}

// RuneIsUpper char
func RuneIsUpper(c rune) bool {
	return 'A' <= c && c <= 'Z'
}
