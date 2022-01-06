package util

// IsLetter returns true if the character is a letter
func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// IsDigit returns true if the character is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// IsWhitespace returns true if the character is whitespace
func IsWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// IsOperator returns true if the character is an operator
func IsOperator(ch byte) bool {
	return ch == '!' || ch == '?' || ch == '[' || ch == ']' || ch == '|'
}
