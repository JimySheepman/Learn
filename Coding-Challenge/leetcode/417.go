var offset = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func pacificAtlantic(matrix [][]int) [][]int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return nil
    }
    m := len(matrix)
    n := len(matrix[0])
    
    resP := make([][]bool, m)
    for i := 0; i < m; i++ {
        resP[i] = make([]bool, n)
    }
    
    for i := 0; i < n; i++ {
        resP[0][i] = true
    }
    
    for i := 0; i < m; i++ {
        resP[i][0] = true
    }
    
    for i := 0; i < n; i++ {
        flood(matrix, m, n, 0, i, resP)
    }
    
    for i := 1; i < m; i++ {
        flood(matrix, m, n, i, 0, resP)
    }
    
    resA:= make([][]bool, m)
    for i := 0; i < m; i++ {
        resA[i] = make([]bool, n)
    }
    
    for i := 0; i < n; i++ {
        resA[m - 1][i] = true
    }
    
    for i := 0; i < m; i++ {
        resA[i][n - 1] = true
    }
    
    for i := 0; i < n; i++ {
        flood(matrix, m, n, m - 1, i, resA)
    }
    
    for i := 0; i < m; i++ {
        flood(matrix, m, n, i, n - 1, resA)
    }
    
    var res [][]int
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if resP[i][j] && resA[i][j] {
                res = append(res, []int{i, j})
            }
        }
    }
    return res
}

func flood(matrix [][]int, m int, n int, i int, j int, res [][]bool) {
    for _, v := range offset {
        i2 := i + v[0]
        j2 := j + v[1]
        if i2 < 0 || i2 >= m || j2 < 0 || j2 >= n {
            continue
        }
        if res[i2][j2] {
            continue
        }
        if matrix[i2][j2] < matrix[i][j] {
            continue 
        }
        res[i2][j2] = true
        flood(matrix, m, n, i2, j2, res)
    }
}
