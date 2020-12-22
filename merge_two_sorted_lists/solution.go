package merge_two_sorted_lists


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    // This one is just the dummy node, doesn't really matter other than 
    // to use to refer to the actual head
    placeholder := &ListNode{}
    merged := placeholder
    
    // While there are still nodes
    for nil != l1 && nil != l2 {
        // Take either l1 or l2 to be pointed to next in the merged LL
        // in ascending order
        if l1.Val <= l2.Val {
            merged.Next = l1
            l1 = l1.Next
        } else {
            merged.Next = l2
            l2 = l2.Next
        }
        
        
        // Move to the newly added node as the current node
        merged = merged.Next
    }
    
    // Exit condition for loop was met, stick the remaining node on the end
    if nil == l1 {
        merged.Next = l2
    } else {
        merged.Next = l1
    }
    
    return placeholder.Next
}
