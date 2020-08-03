package review

import "regexp"

func sanitizeERB(data []byte) []byte {
	var replacer = regexp.MustCompile(`\<\%\=?[^\>]+?\%\>`)
	return replacer.ReplaceAll(data, []byte{})
}
