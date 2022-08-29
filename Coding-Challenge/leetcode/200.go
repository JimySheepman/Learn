func numIslands(grid [][]byte) int {
    if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    result := 0
    m := len(grid)
    n := len(grid[0])
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                dfs(&grid, i, j)
                result++
            }
        }
    }
    return result
}

func dfs(grid *[][]byte, i int, j int) {
    (*grid)[i][j] = '0'
    if i - 1 >= 0 && (*grid)[i-1][j] == '1' {
        dfs(grid, i - 1, j)
    }
    if j + 1 < len((*grid)[0]) && (*grid)[i][j+1] == '1' {
        dfs(grid, i, j + 1)
    }
    if j - 1 >= 0 && (*grid)[i][j-1] == '1' {
        dfs(grid, i, j - 1)
    }
    if i + 1 < len(*grid) && (*grid)[i+1][j] == '1' {
        dfs(grid, i + 1, j)
    }
}
