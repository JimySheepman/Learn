var resp [][]string
func partition(s string) [][]string {
    resp = [][]string{}
    backtracking(s,[]string{})
    return resp
}

func backtracking(s string ,temp []string){
    if len(s) == 0 {
        subResp := make([]string , len(temp))
        copy(subResp,temp)
        resp = append(resp,subResp)
        return
    }
    for i := 0 ; i<=len(s) ; i++{
        if isPalindrome(s[0:i]){
            backtracking(s[i:] , append(temp , s[0:i]))
        }
    }
}


func isPalindrome(s string) bool {
    if len(s) == 0{
        return false
    }
    left := 0
    right := len(s) -1
    for left < right{
        if s[left] == s[right]{
            left++
            right--
        } else {
            return false
        }
    }
    return true;
}
