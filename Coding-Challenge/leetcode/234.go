/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    p1, p2 := head, head
    for p2.Next != nil && p2.Next.Next != nil {
        p1 = p1.Next
        p2 = p2.Next.Next
    }
    
    var front *ListNode
    mid, end := p1.Next, p1.Next
    p1.Next = nil
    p1 = mid
    for mid != nil {
        end = mid.Next
        mid.Next = front
        front, mid = mid, end
    }
    
    for front != nil {
        if head.Val != front.Val {
            return false
        }
        head = head.Next
        front = front.Next
    }
    return true
}

