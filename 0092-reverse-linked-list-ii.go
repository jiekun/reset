// Given the head of a singly linked list and two integers left and right where left <= right, reverse the nodes of the list from position left to position right, and return the reversed list.

// Example 1:

// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]
// Example 2:

// Input: head = [5], left = 1, right = 1
// Output: [5]

// Constraints:

// The number of nodes in the list is n.
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n

// Follow up: Could you do it in one pass?

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    node := head
    var prev, last *ListNode
    i := 1
    for ; i <= left; i++ {
        last = prev
        prev = node
        node = node.Next
    }
    // don't touch last
    // find next
    //reverse current
    for ; i <= right; i++ {
        tmp := node.Next
        node.Next = prev
        prev = node
        node = tmp
    }

    if last == nil {
        head.Next = node
        head = prev
        return head
    }

    if last.Next != nil {
        last.Next.Next = node
        last.Next = prev
    }

    return head
}