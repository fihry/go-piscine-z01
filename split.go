package piscine

func SplitWhiteSpaces(s string) []string {
    res := []string{}
    word := ""
    for j, i := range s {
        if i != ' ' && i != '\n' && i != '\t' {
            word += string(i)
        } else if ((i == ' ' || i == '\t' || i == '\n') && word != "") || j == len(s)-1 {
            res = append(res, word)
            word = ""
        }
    }
    if word != "" {
        res = append(res, word)
    }
    return res
}

func Index(s string, toFind string) int {
    for i := 0; i <= len(s)-len(toFind); i++ {
        if s[i:i+len(toFind)] == toFind {
            return i
        }
    }
    return -1
}

func Split(s, sep string) []string {
    for i := 0; i < len(s)-1; i++ {
        if Index(s, sep) != -1 {
            s = s[0:Index(s, sep)] + " " + s[Index(s, sep)+len(sep):]
        }
    }
    return SplitWhiteSpaces(s)
}