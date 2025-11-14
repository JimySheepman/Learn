func findJudge(n int, trust [][]int) int {
    cantBeJudge := make([]bool, n)
    countOfTrusters := make([]int, n)

    for _, t := range trust {
        cantBeJudge[t[0]-1] = true  
        countOfTrusters[t[1]-1]++
    }

    for i, cant := range cantBeJudge {
        if !cant && countOfTrusters[i] == n - 1 {
            return i+1
        }
    }
    return -1
}
