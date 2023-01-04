package urlutil

import (
	"net/url"
	"strings"
)

func UrlEncodeByMap(m map[string]string) string {
	if len(m) == 0 {
		return ""
	}
	param := url.Values{}
	for k, v := range m {
		param.Add(k, v)
	}
	unescape, err := url.QueryUnescape(param.Encode())
	if err != nil {
		return ""
	}
	return unescape
}

func UrlDecodeToMap(str string) map[string]string {
	values, err := url.ParseQuery(str)
	if err != nil {
		return nil
	}
	m := make(map[string]string)
	for k, v := range values {
		m[k] = v[0]
	}
	return m
}

// UrlEncode 编码 url 字符串。
func UrlEncode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // escape query data
		return s[0:pos+1] + url.QueryEscape(s[pos+1:])
	}
	return s
}

// UrlDecode 解码 url 字符串。
func UrlDecode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // un-escape query data
		qy, err := url.QueryUnescape(s[pos+1:])
		if err == nil {
			return s[0:pos+1] + qy
		}
	}
	return s
}

func RawUrlEncode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}
func RawUrlDecode(str string) (string, error) {
	return url.QueryUnescape(strings.Replace(str, "%20", "+", -1))
}
