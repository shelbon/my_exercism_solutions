package parsinglogfiles
import ("regexp"
        "fmt")
func IsValidLine(text string) bool {
    re,err:= regexp.Compile(`^(\[TRC\]|\[WRN\]|\[INF\]|\[ERR\]|\[DBG\]|\[FTL\])`)
    if err != nil {
        panic("Malformated regex")
    }
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	  re,err:= regexp.Compile(`<(~|\*|=|-)*>`)
    if err != nil {
        panic("Malformated regex")
    }
	return re.Split(text,-1)
}

func CountQuotedPasswords(lines []string) int {
	re,err:= regexp.Compile(`(?iU)".*password.*"`)
    quotedPasswordsCount:=0
    if err != nil {
        panic("Malformated regex")
    }
    for _,line := range lines {
        matches:=re.FindStringSubmatch(line)
        if matches != nil {
            quotedPasswordsCount+=len(matches);
        }
    }
	return quotedPasswordsCount
}

func RemoveEndOfLineText(text string) string {
	re,err:= regexp.Compile(`end-of-line\d+`)
    if err != nil {
        panic("Malformated regex")
    }
	return re.ReplaceAllString(text,"");
}

func TagWithUserName(lines []string) []string {
    tagWithUsername := make([]string, 0, len(lines))
    prefix:="[USR]"
    re,err:= regexp.Compile(`User\s+(\w*)`)
    if err!= nil {
        panic("Malformated regex")
    }
    for _,line := range lines {
           matches:=re.FindStringSubmatch(line)
           if len(matches)> 0 {
            tagWithUsername = append(tagWithUsername, fmt.Sprintf("%s %s %s",prefix,matches[1],line))
        }else{
            tagWithUsername = append(tagWithUsername,line)
        }

       }
	return tagWithUsername
}
