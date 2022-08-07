package helpers

import "regexp"

func CheckFormatFuncCallback(payload string) (match bool) {
	match, _ = regexp.MatchString("[a-z] [0-9]", payload)
	if match {
		return match
	}
	match, _ = regexp.MatchString("0-9]", payload)
	return match
}
