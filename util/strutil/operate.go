package strutil

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// StrToLower 转换成小写字母
func StrToLower(str string) string {
	runeArr := []rune(str)
	for i := range runeArr {
		if runeArr[i] >= 65 && runeArr[i] <= 90 {
			runeArr[i] += 32
		}
	}
	return string(runeArr)
}

// ConcatString 连接字符串,性能比fmt.Sprintf和+号要好
func ConcatString(s ...string) string {
	if len(s) == 0 {
		return ""
	}
	var buffer bytes.Buffer
	for _, i := range s {
		buffer.WriteString(i)
	}
	return buffer.String()
}

// StringToUint64 字符串转uint64
func StringToUint64(str string) (uint64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return uint64(valInt), nil
}

// StringToInt64 字符串转int64
func StringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return int64(valInt), nil
}

// StringToInt 字符串转int
func StringToInt(str string) (int, error) {
	if str == "" {
		return 0, nil
	}
	valInt, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}

	return valInt, nil
}

// Bytes2String 字节切片转字符串
func Bytes2String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// String2Bytes 字符串转字节切片
func String2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

// SubStr 截取字符串，并返回实际截取的长度和子串
func SubStr(str string, start, end int64) (int64, string, error) {
	reader := strings.NewReader(str)

	// Calling NewSectionReader method with its parameters
	r := io.NewSectionReader(reader, start, end)

	// Calling Copy method with its parameters
	var buf bytes.Buffer
	n, err := io.Copy(&buf, r)
	return n, buf.String(), err
}

// SubstrTarget 在字符串中查找指定子串，并返回left或right部分
func SubstrTarget(str string, target string, turn string, hasPos bool) (string, error) {
	pos := strings.Index(str, target)

	if pos == -1 {
		return "", nil
	}

	if turn == "left" {
		if hasPos {
			pos = pos + 1
		}
		return str[:pos], nil
	} else if turn == "right" {
		if !hasPos {
			pos = pos + 1
		}
		return str[pos:], nil
	} else {
		return "", errors.New("params 3 error")
	}
}

// GetStringUtf8Len 获得字符串按照uft8编码的长度
func GetStringUtf8Len(str string) int {
	return utf8.RuneCountInString(str)
}

// Utf8Index 按照uft8编码匹配子串，返回开头的索引
func Utf8Index(str, substr string) int {
	index := strings.Index(str, substr)
	if index < 0 {
		return -1
	}
	return utf8.RuneCountInString(str[:index])
}

// JoinStringAndOther 连接字符串和其他类型
func JoinStringAndOther(val ...interface{}) string {
	return fmt.Sprint(val...)
}

// UcFirst 首字母大写
func UcFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)
	f := rs[0]

	if 'a' <= f && f <= 'z' {
		return string(unicode.ToUpper(f)) + string(rs[1:])
	}
	return s
}

// LcFirst 首字母小写
func LcFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	rs := []rune(s)
	f := rs[0]

	if 'A' <= f && f <= 'Z' {
		return string(unicode.ToLower(f)) + string(rs[1:])
	}
	return s
}

// FormatPrivateKey 格式化 普通应用秘钥
func FormatPrivateKey(privateKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN RSA PRIVATE KEY-----\n")
	rawLen := 64
	keyLen := len(privateKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(privateKey[start:])
		} else {
			buffer.WriteString(privateKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END RSA PRIVATE KEY-----\n")
	pKey = buffer.String()
	return
}

// FormatPublicKey 格式化 普通支付宝公钥
func FormatPublicKey(publicKey string) (pKey string) {
	var buffer strings.Builder
	buffer.WriteString("-----BEGIN PUBLIC KEY-----\n")
	rawLen := 64
	keyLen := len(publicKey)
	raws := keyLen / rawLen
	temp := keyLen % rawLen
	if temp > 0 {
		raws++
	}
	start := 0
	end := start + rawLen
	for i := 0; i < raws; i++ {
		if i == raws-1 {
			buffer.WriteString(publicKey[start:])
		} else {
			buffer.WriteString(publicKey[start:end])
		}
		buffer.WriteByte('\n')
		start += rawLen
		end = start + rawLen
	}
	buffer.WriteString("-----END PUBLIC KEY-----\n")
	pKey = buffer.String()
	return
}

// CamelCase 将字符串转换为驼峰式字符串。
func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""
	blankSpace := " "
	regex, _ := regexp.Compile("[-_&]+")
	ss := regex.ReplaceAllString(s, blankSpace)
	for i, v := range strings.Split(ss, blankSpace) {
		vv := []rune(v)
		if i == 0 {
			if vv[i] >= 65 && vv[i] <= 96 {
				vv[0] += 32
			}
			result += string(vv)
		} else {
			result += Capitalize(v)
		}
	}

	return result
}

// Capitalize 将字符串的第一个字符转换为大写，其余字符转换为小写。
func Capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}

	out := make([]rune, len(s))
	for i, v := range s {
		if i == 0 {
			out[i] = unicode.ToUpper(v)
		} else {
			out[i] = unicode.ToLower(v)
		}
	}

	return string(out)
}

// KebabCase 将字符串转换为 kebab-case
func KebabCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "-")
}

// SnakeCase 将字符串转换为蛇形大小写
func SnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	regex := regexp.MustCompile(`[\W|_]+`)
	blankSpace := " "
	match := regex.ReplaceAllString(s, blankSpace)
	rs := strings.Split(match, blankSpace)

	var result []string
	for _, v := range rs {
		splitWords := splitWordsToLower(v)
		if len(splitWords) > 0 {
			result = append(result, splitWords...)
		}
	}

	return strings.Join(result, "_")
}

// splitWordsToLower 将将字符串拆分为多个字符串,并转小写
func splitWordsToLower(s string) []string {
	var result []string

	upperIndexes := upperIndex(s)
	l := len(upperIndexes)
	if upperIndexes == nil || l == 0 {
		if s != "" {
			result = append(result, s)
		}
		return result
	}
	for i := 0; i < l; i++ {
		if i < l-1 {
			result = append(result, strings.ToLower(s[upperIndexes[i]:upperIndexes[i+1]]))
		} else {
			result = append(result, strings.ToLower(s[upperIndexes[i]:]))
		}
	}
	return result
}

// upperIndex 得到一个 int 切片，其中元素都是字符串的大写 char 索引
func upperIndex(s string) []int {
	var result []int
	for i := 0; i < len(s); i++ {
		if 64 < s[i] && s[i] < 91 {
			result = append(result, i)
		}
	}
	if len(s) > 0 && result != nil && result[0] != 0 {
		result = append([]int{0}, result...)
	}

	return result
}

// Reverse 返回字符顺序与给定字符串相反的字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Quote 返回双引号的字符串
func Quote(s string) string {
	return strconv.Quote(s)
}

// AddSlashes 为字符串添加斜线。
func AddSlashes(s string) string {
	if ln := len(s); ln == 0 {
		return ""
	}

	var buf bytes.Buffer
	for _, char := range s {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}

	return buf.String()
}

// StripSlashes 去除字符串的斜杠。
func StripSlashes(s string) string {
	ln := len(s)
	if ln == 0 {
		return ""
	}

	var skip bool
	var buf bytes.Buffer

	for i, char := range s {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < ln && s[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}

	return buf.String()
}

// Trim  如果 cutSet 为空，将去除空格。
func Trim(s string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.Trim(s, cutSet[0])
		}

		return strings.Trim(s, strings.Join(cutSet, ""))
	}

	return strings.TrimSpace(s)
}

// LTrim 字符串中的字符。如果 cutSet 为空，将去除空格。
func LTrim(s string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.TrimLeft(s, cutSet[0])
		}

		return strings.TrimLeft(s, strings.Join(cutSet, ""))
	}

	return strings.TrimLeft(s, " ")
}

// RTrim 字符串中的字符。如果 cutSet 为空，将去除空格。
func RTrim(s string, cutSet ...string) string {
	if ln := len(cutSet); ln > 0 && cutSet[0] != "" {
		if ln == 1 {
			return strings.TrimRight(s, cutSet[0])
		}
		return strings.TrimRight(s, strings.Join(cutSet, ""))
	}

	return strings.TrimRight(s, " ")
}

// UpperEnglishWord 将每个单词的第一个字符更改为大写
func UpperEnglishWord(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	inWord := true
	buf := make([]byte, 0, len(s))
	i := 0
	rs := []rune(s)
	if RuneIsLower(rs[i]) {
		buf = append(buf, []byte(string(unicode.ToUpper(rs[i])))...)
	} else {
		buf = append(buf, []byte(string(rs[i]))...)
	}
	for j := i + 1; j < len(rs); j++ {
		if !RuneIsWord(rs[i]) && RuneIsWord(rs[j]) {
			inWord = false
		}
		if RuneIsLower(rs[j]) && !inWord {
			buf = append(buf, []byte(string(unicode.ToUpper(rs[j])))...)
			inWord = true
		} else {
			buf = append(buf, []byte(string(rs[j]))...)
		}
		if RuneIsWord(rs[j]) {
			inWord = true
		}
		i++
	}
	return string(buf)
}

// PosFlag type
type PosFlag uint8

// Position for padding/resize string
const (
	PosLeft PosFlag = iota
	PosRight
	PosMiddle
)

// Padding 填充字符串。
func Padding(s, pad string, length int, pos PosFlag) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if pad == "" || pad == " " {
		mark := ""
		if pos == PosRight { // to right
			mark = "-"
		}

		// padding left: "%7s", padding right: "%-7s"
		tpl := fmt.Sprintf("%s%d", mark, length)
		return fmt.Sprintf(`%`+tpl+`s`, s)
	}

	if pos == PosRight { // to right
		return s + Repeat(pad, -diff)
	}
	return Repeat(pad, -diff) + s
}

// PadLeft 左边填充一个字符串
func PadLeft(s, pad string, length int) string {
	return Padding(s, pad, length, PosLeft)
}

// PadRight 右边填充一个字符串
func PadRight(s, pad string, length int) string {
	return Padding(s, pad, length, PosRight)
}

// Repeat 重复一个字符串
func Repeat(s string, times int) string {
	if times <= 0 {
		return ""
	}
	if times == 1 {
		return s
	}

	ss := make([]string, 0, times)
	for i := 0; i < times; i++ {
		ss = append(ss, s)
	}

	return strings.Join(ss, "")
}

// Resize 按给定的长度和对齐设置调整字符串的大小。填充空间。
func Resize(s string, length int, align PosFlag) string {
	diff := len(s) - length
	if diff >= 0 { // do not need padding.
		return s
	}

	if align == PosMiddle {
		strLn := len(s)
		padLn := (length - strLn) / 2
		padStr := string(make([]byte, padLn))

		if diff := length - padLn*2; diff > 0 {
			s += " "
		}
		return padStr + s + padStr
	}

	return Padding(s, " ", length, align)
}
