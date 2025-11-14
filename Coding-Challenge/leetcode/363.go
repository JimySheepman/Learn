func maxSumSubmatrix(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    maxSum := math.MinInt32

    for i:=0; i<m; i++ {
        for j:=1;j<n; j++{
            matrix[i][j] += matrix[i][j-1]
        }
    }
    for l:=0; l<n; l++ {
        for r:=l; r < n; r++ {
            prefixSum, set := 0, Set{0}
            for i:=0; i< m; i++ {
                sum := matrix[i][r]
                if l > 0 {
                    sum -= matrix[i][l-1]
                }
                prefixSum += sum
                idx := sort.SearchInts(set, prefixSum-k)
                if idx != len(set) {
                    maxSum = max(maxSum, prefixSum-set[idx])
                }
                set.Insert(prefixSum)
            }
        }
    }
    
    return maxSum
}

type Set []int

func (s *Set) Insert(x int) {
	i := sort.SearchInts(*s, x)
	if i == len(*s) {
		*s = append(*s, x)
	} else if (*s)[i] != x {
		*s = append(*s, 0)
		copy((*s)[i+1:], (*s)[i:])
		(*s)[i] = x
	}
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
