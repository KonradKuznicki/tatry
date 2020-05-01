package lib

import "regexp"

var re = regexp.MustCompile(`[^a-zA-Z0-9\-\.]+`)

func SafeFileName(fileName string) string {
	return re.ReplaceAllString(fileName, `_`)
}
