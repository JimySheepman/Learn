/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {
    col:= 0
    minCol:= 0
    maxCol:= 0

    vert:=make(map[int][]int)
    hor:=make(map[int][]int)
    q:=[]*TreeNode{root}
    c:=[]int{col}
    cur:=root
    for len(q)>0{
        n:=len(q)
        for i:=0;i<n;i++{
            cur,q = q[0],q[1:]
            col,c = c[0],c[1:]
            if col<minCol {
                minCol = col
            }
            if col>maxCol {
                maxCol = col
            }
            curCol:=col
            hor[col] = append(hor[col],cur.Val)
            if cur.Left!=nil {
                q = append(q,cur.Left)
                c = append(c,curCol-1)
            }
            if cur.Right!=nil {
                q = append(q,cur.Right)
                c = append(c,curCol+1)
            }
        }
        for i:=minCol;i<=maxCol;i++{
            sort.Ints(hor[i])
            vert[i] = append(vert[i],hor[i]...)
            hor[i] = []int{}
        }        
    }
    ans:=[][]int{}
    for i:=minCol;i<=maxCol;i++{
        ans = append(ans,vert[i])
    }
    return ans
}
