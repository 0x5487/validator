package validator

import "regexp"

var (
	_regExp *regexp.Regexp
)

func init() {
	_regExp, _ = regexp.Compile("^[a-zA-Z0-9]*$")
}

func IsAlphaNumeric(text string) bool {

	return _regExp.MatchString(text)
}
