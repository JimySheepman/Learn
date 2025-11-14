func diagonalSort(mat [][]int) [][]int {
    m := len(mat)
    n := len(mat[0])
    tmp := make([][]int, m*n)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            diff := j - i
            if diff >= 0 {
                tmp[diff] = append(tmp[diff], mat[i][j])
            } else {
                diff = diff * -n 
                tmp[diff] = append(tmp[diff], mat[i][j])
            }
        }
    }
    for i := range tmp {
        sort.Slice(tmp[i], func(a, b int) bool {
            return tmp[i][a] < tmp[i][b]
        })
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            diff := j - i
            if diff >= 0 {
                mat[i][j] = tmp[diff][0]
                tmp[diff] = tmp[diff][1:]
            } else {
                diff = diff * -n
                mat[i][j] = tmp[diff][0]
                tmp[diff] = tmp[diff][1:]
            }
        }
    }
    return mat
}
