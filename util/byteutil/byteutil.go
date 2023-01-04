package byteutil

// IsLetterUpper 检查给定字节 b 是否为大写。
func IsLetterUpper(b byte) bool {
	if b >= byte('A') && b <= byte('Z') {
		return true
	}
	return false
}

// IsLetterLower 检查给定字节 b 是否为小写。
func IsLetterLower(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	}
	return false
}

// IsLetter 检查给定的字节 b 是否是一个字母。
func IsLetter(b byte) bool {
	return IsLetterUpper(b) || IsLetterLower(b)
}
