package remove_nth_node_from_end_of_list


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // 0 indicates to chop off the head beacause the entire LL was traversed
    if 0 == depth_first_search(head, n) {
        head = head.Next
    }
    
    return head
}


func depth_first_search(node *ListNode, n int) int {
    // Exit condition no node
    if nil == node {
        return n
    }
    
    n = depth_first_search(node.Next, n)
    
    // Chop the node in front of the current node out
    if 0 == n {
        node.Next = node.Next.Next
    }
    
    // Decrement per stack frame represents a traversal
    return n - 1
}
