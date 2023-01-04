package strutil

import "strings"

// StrPos strings.Index 的别名
// 返回 s 中 substr 的第一个实例的索引，如果 substr 不存在于 s 中，则返回 -1。
func StrPos(s, sub string) int {
	return strings.Index(s, sub)
}
