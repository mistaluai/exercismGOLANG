package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re, _ := regexp.Compile(`TRC|DBG|INF|WRN|ERR|FTL`)
	if len(text) >= 5 {
		text = text[0:5]
	} else {
		return false
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, _ := regexp.Compile(`<(~|\*|=|-)*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, _ := regexp.Compile(`".*[pP][aA][sS][sS][wW][oO][rR][dD].*"`)
	valid := 0
	for _, line := range lines {
		if (re.MatchString(line)) == true {
			valid++
		}
	}
	return valid
}

func RemoveEndOfLineText(text string) string {
	re, _ := regexp.Compile(`(end-of-line\d*)`)
	text = re.ReplaceAllString(text, "")
	return text
}

func TagWithUserName(lines []string) []string {
	re, _ := regexp.Compile(`(User)\s+\w*\d*`)
	reSplitter, _ := regexp.Compile(`\s+`)
	var newLines []string = make([]string, len(lines))
	for i, line := range lines {
		if re.MatchString(line) {
			splittedExpr := reSplitter.Split(re.FindString(line), 2)
			username := splittedExpr[1]
			usr := "[USR] "
			usr = usr + username + " " + line
			line = usr
		}
		newLines[i] = line
	}
	return newLines
}
