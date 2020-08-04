package review

import (
	"regexp"
)

func sanitizeERB(data []byte) []byte {
	var replacer = regexp.MustCompile(`\<\%\=?[^\>]+?\%\>`)
	var replacerVal = regexp.MustCompile(`=(["'])\<\%\=?[^\>]+?\%\>(["'])`)

	dataSanitized := replacerVal.ReplaceAllFunc(data, func(exp []byte) []byte {
		expParts := replacerVal.FindSubmatch(exp)
		if len(expParts) == 3 {
			var result []byte
			result = append(result, byte('='))
			result = append(result, expParts[1][0])
			result = append(result, []byte("MILO WAS HERE!")...)
			result = append(result, expParts[2][0])
			return result
		}
		return exp
	})

	return replacer.ReplaceAll(dataSanitized, []byte{})
}
