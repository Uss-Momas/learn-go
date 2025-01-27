package main

/**
* NOTE: in ASCII TABLE digits are between: 48 and 57 (included)
 */
func _isdigit(c int) int {
	// it is a digit
	if c >= 48 && c <= 57 {
		return 1
	}
	return 0
}
