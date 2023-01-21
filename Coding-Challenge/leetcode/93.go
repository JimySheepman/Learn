func restoreIpAddresses(s string) []string {
    result := []string{}
    dfs(s, "", 0, 4, &result)
    return result
}

func dfs(s, prev string, start, count int, result *[]string) {
    n := len(s)-start
    if n < count || n > count*3 {
        return
    }
    if count == 1 && isValidStr(s[start:]) {
        *result = append(*result, prev+s[start:])
        return
    }
    k := min(n, 3)
    for i := 1; i <= k; i++ {
        nextStart := start+i
        if !isValidStr(s[start:nextStart]) {
            continue
        }
        dfs(s, prev+s[start:nextStart]+".", nextStart, count-1, result)
    }
}

func isValidStr(s string) bool {
    if len(s) == 1 {
        return true
    }
    if s[0] == '0' {
        return false
    }
    tmp := 0
    for i := range s {
        tmp = tmp*10 + int(s[i]-'0')
    }
    return tmp <= 255
}

func min(a,b int) int {
    if a < b {
        return a
    }
    return b
}

