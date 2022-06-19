package codes

import "regexp"

//	Regular Expression for ANSI Escape Codes
const REGEXP = "[\u001B\u009B][[\\]()#;?]*(?:(?:(?:(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]+)*|[a-zA-Z\\d]+(?:;[-a-zA-Z\\d\\/#&.:=?%@~_]*)*)?\u0007)" + "|" + "(?:(?:\\d{1,4}(?:;\\d{0,4})*)?[\\dA-PR-TZcf-nq-uy=><~]))"

//	Regular Expression for ANSI Escape Codes
var REGEX = regexp.MustCompile(REGEXP)

//	Strip ANSI Codes from string
func Strip(str string) string {
	return REGEX.ReplaceAllString(str, "")
}
