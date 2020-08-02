package review

import "regexp"

func sanitizeERB(data []byte) []byte {
	var replacer = regexp.MustCompile(`<%[^%]*%>`)
	content := replacer.ReplaceAllString(string(data), "ERB")

	return []byte(content)
}
